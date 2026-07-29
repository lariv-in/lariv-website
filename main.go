package main

import (
	"log/slog"

	"github.com/lariv-in/lariv"
	"github.com/lariv-in/lariv/plugins/p_blog"
	"github.com/lariv-in/lariv/plugins/p_dashboard"
	"github.com/lariv-in/lariv/plugins/p_filesystem"
	"github.com/lariv-in/lariv/plugins/p_no_signup"
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
		p_no_signup.GetPlugin(),
		p_blog.GetPlugin(),
	}

	app, err := lariv.NewBuilder().AddPlugins(plugins).LoadConfigFromFile("config.toml")
	if err != nil {
		panic(err)
	}

	if err := app.Start(); err != nil {
		slog.Error(err.Error())
	}
}
