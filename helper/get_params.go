package helper

import (
	"net/http"
	"strings"
)

func GetParamURL(req *http.Request,prefix string) string {
	path := strings.TrimPrefix(req.URL.Path,prefix);

	if strings.Contains(path,"/"){
		parts := strings.Split(path, "/");

		return parts[0];
	}

	return path;
}