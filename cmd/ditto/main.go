package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"

	"github.com/arvingarciabtw/ditto/internal/config"
	"github.com/arvingarciabtw/ditto/internal/evdev"
	"github.com/arvingarciabtw/ditto/internal/tui"
)

func main() {
	cfg := config.LoadWithFlags()

	devs, err := evdev.Devices()
	if err != nil {
		evdev.PrintDeviceError(err)
		os.Exit(1)
	}

	p := tea.NewProgram(tui.InitModel(cfg))
	for _, dev := range devs {
		go evdev.ListenToKeyboard(p, dev)
	}
	_, err = p.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
