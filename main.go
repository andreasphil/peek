package main

import (
	"context"
	"embed"
	"net/http"
	"os"
	"os/signal"

	"github.com/andreasphil/peek/cli"
	"github.com/andreasphil/peek/lib"
	"github.com/andreasphil/peek/server"
	"github.com/charmbracelet/log"
)

//go:embed static
var static embed.FS

func serve(config *cli.Config) {
	context, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	server := server.NewServer(server.ServerInit{
		Filename:    config.Filename,
		FilePreview: lib.NewPreviewService(config.AllowUnsafe),
		Port:        config.Port,
		Static:      static,
	})

	go func() {
		log.Infof("serving preview of %v at http://localhost:%v", config.Filename, config.Port)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("server exited with an error", "err", err)
			os.Exit(1)
		}
	}()

	<-context.Done()
	server.Shutdown(context)
}

func main() {
	serve(cli.ParseFlags())
}
