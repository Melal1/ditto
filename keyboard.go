package main

import (
	"fmt"

	evdev "github.com/gvalkov/golang-evdev"
)

type globalKeyMsg struct {
	code uint16
	down bool
}

func device() (*evdev.InputDevice, error) {
	maxEvents := 32
	for e := range maxEvents {
		device, err := evdev.Open(fmt.Sprintf("/dev/input/event%d", e))
		if err != nil {
			continue
		}
		if isKeyboardDevice(device) {
			return device, nil
		}
		_ = device.File.Close()
	}

	return nil, fmt.Errorf("no keyboard device found")
}

func isKeyboardDevice(device *evdev.InputDevice) bool {
	for cpbType, cpbCodes := range device.Capabilities {
		if cpbType.Type != evdev.EV_KEY {
			continue
		}
		for _, cpbCode := range cpbCodes {
			if cpbCode.Code == evdev.KEY_A {
				return true
			}
		}
	}
	return false
}
