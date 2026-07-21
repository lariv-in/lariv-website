package main

import (
	"log/slog"

	"github.com/lariv-in/lariv"
	"github.com/lariv-in/lariv/plugins/p_dashboard"
	"github.com/lariv-in/lariv/plugins/p_filesystem"
	"github.com/lariv-in/lariv/plugins/p_users"
	"github.com/lariv-in/lariv/plugins/p_website"
	"github.com/lariv-in/lariv/registry"
)

func main() {
	plugins := []registry.Pair[string, lariv.Plugin]{
		p_dashboard.GetPlugin(),
		p_filesystem.GetPlugin(),
		p_users.GetPlugin(),
		p_website.GetPlugin(),
	}

	config, err := lariv.LoadConfigFromFile("config.toml", plugins)
	if err != nil {
		panic(err)
	}

	if err := lariv.Start(config, plugins); err != nil {
		slog.Error(err.Error())
	}
}
