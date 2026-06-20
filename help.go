package main

import (
	bkey "charm.land/bubbles/v2/key"
)

type bindings struct {
	Layout  bkey.Binding
	Size    bkey.Binding
	HideKey bkey.Binding
	Help    bkey.Binding
	Quit    bkey.Binding
}

func (c bindings) ShortHelp() []bkey.Binding {
	return []bkey.Binding{c.Help, c.Quit}
}

func (c bindings) FullHelp() [][]bkey.Binding {
	return [][]bkey.Binding{
		{c.Size},
		{c.Layout},
		{c.HideKey},
	}
}

var commands = bindings{
	Layout: bkey.NewBinding(
		bkey.WithKeys("ctrl+shift+l"),
		bkey.WithHelp("^l", "layout"),
	),
	Size: bkey.NewBinding(
		bkey.WithKeys("ctrl+shift+s"),
		bkey.WithHelp("^s", "size"),
	),
	HideKey: bkey.NewBinding(
		bkey.WithKeys("ctrl+shift+h"),
		bkey.WithHelp("^h", "hide"),
	),
	Help: bkey.NewBinding(
		bkey.WithKeys("?"),
		bkey.WithHelp("?", "help"),
	),
	Quit: bkey.NewBinding(
		bkey.WithKeys("q", "ctrl+c"),
		bkey.WithHelp("q", "quit"),
	),
}
