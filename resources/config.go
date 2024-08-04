package resources

import (
	"embed"
	"net/http"
)

//go:embed templates/*.html
var Templates embed.FS

//go:embed assets
var assets embed.FS

func FileServer(mux *http.ServeMux) {
    fileServer := http.FileServer(http.FS(assets));
    mux.Handle("/assets/", fileServer);
}