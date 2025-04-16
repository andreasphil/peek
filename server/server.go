package server

import (
	"fmt"
	"io/fs"
	"net/http"

	"github.com/andreasphil/peek/adapters"
	"github.com/charmbracelet/log"
)

type ServerInit struct {
	FilePreview adapters.FilePreviwer
	Static      fs.FS
	Port        string
	Filename    string
}

func NewServer(init ServerInit) http.Server {
	router := http.NewServeMux()

	router.HandleFunc("/", handlePreview(init.FilePreview, init.Filename))
	router.Handle("/static/", http.FileServerFS(init.Static))

	return http.Server{
		Addr:    fmt.Sprintf(":%v", init.Port),
		Handler: router,
	}
}

func handlePreview(filePreview adapters.FilePreviwer, filename string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		preview, err := filePreview.ForFile(filename)

		if err != nil {
			log.Error("could not render file contents", "filename", filename, "err", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(preview))
	}
}
