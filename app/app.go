package app

import (
	"context"
	"fmt"
	"github.com/Yiivgeny/incotex-mercury-client/client"
	"github.com/Yiivgeny/incotex-mercury-client/client/methods/read_parameter"
	"github.com/Yiivgeny/incotex-mercury-client/protocol"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"mercury-client/helper"
	"sync"
	"time"
)

const Version = "undefined"
const Name = "Mercury Client " + Version

var Logger *zap.Logger

func Init() error {
	for _, plugin := range plugins {
		err := plugin.Init()
		if err != nil {
			return errors.WithMessagef(err, "initialize %s", plugin.Name())
		}
	}

	return nil
}

func Finish() {
	for _, plugin := range plugins {
		plugin.Finish()
	}
}

func Serve(ctx context.Context) error {

	address := protocol.Address(0)

	cfg := client.NewConfig(9600, 2)
	cfg.Host = "192.168.90.231:8899"
	cfg.ResponseTimeout += time.Second * 10
	auth := client.Auth{
		AccessLevel: protocol.AccessLevelUser,
		Password:    []byte{1, 1, 1, 1, 1, 1},
	}

	transport, err := client.NewTransportTCP(&cfg)
	if err != nil {
		return err
	}
	defer func() {
		_ = transport.Close()
	}()

	c := client.NewClient(&cfg, transport)
	if err = c.TestCommunication(address); err != nil {
		return err
	}

	if err = c.OpenChannel(address, auth); err != nil {
		return err
	}

	defer func() {
		_ = c.CloseChannel(address)
	}()

	request, individual := read_parameter.NewIndividualOptions()
	if err = c.Request(address, request, individual); err != nil {
		return err
	}

	for _, plugin := range plugins {
		if err := plugin.RegisterDevice(individual); err != nil {
			Logger.Error(fmt.Sprintf("register device with plugin %s", plugin.Name()), zap.Error(err))
			// TODO: filter plugins
		}
	}

	wg := sync.WaitGroup{}

	req, instantIndicators := read_parameter.NewInstantIndicators()
	for _, plugin := range plugins {
		if err := plugin.RegisterInstantIndicators(individual, instantIndicators); err != nil {
			Logger.Warn(fmt.Sprintf("register instant indicator for plugin %s", plugin.Name()), zap.Error(err))
		}
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := helper.Schedule(ctx, func(ctx context.Context) error {
			if err := c.Request(address, req, instantIndicators); err != nil {
				return err
			}
			for _, plugin := range plugins {
				if err := plugin.PushInstantIndicators(individual, instantIndicators); err != nil {
					Logger.Warn(fmt.Sprintf("push instant indicator for plugin %s", plugin.Name()), zap.Error(err))
				}
			}
			return nil
		}, time.Minute)
		if err != nil {
			Logger.Error("process instant indicators", zap.Error(err))
		}
	}()

	wg.Wait()
	return nil
}
