package main

import (
	"context"
	"embed"
	"io/fs"

	"gitlab.as203038.net/AS203038/looking-glass/server"
)

//go:embed all:dist
var webemned embed.FS

func main() {
	web, err := fs.Sub(webemned, "dist")
	if err != nil {
		panic(err)
	}
	server.Start(context.Background(), web)
}
