package app

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"runtime"
	"strings"
)

func bindHttpHandlers() {
	fileServer := http.NewServeMux()
	fileServer.Handle("/", useCustom404(http.FileServer(http.Dir("pub/"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if p := recover(); p != nil {
				buf := make([]byte, 1<<16)
				runtime.Stack(buf, false)
				reason := fmt.Sprintf("%v: %s", r, buf)
				appLog.Critical("Runtime failure, reason -> %s", reason)
			}
		}()
		switch {
		case strings.Contains(r.URL.Path, "."):
			fileServer.ServeHTTP(w, r)
			return
		case r.URL.Path != "/":
			write404(w)
			return
		}
		pageHandler(w, r)
	})
}

func write404(w http.ResponseWriter) {
	w.Write([]byte("gala not found!"))
}

type hijack404 struct {
	http.ResponseWriter
}

func (h *hijack404) WriteHeader(code int) {
	if code == 404 {
		write404(h.ResponseWriter)
		panic(h)
	}
	h.ResponseWriter.WriteHeader(code)
}

func useCustom404(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hijack := &hijack404{w}
		defer func() {
			if p := recover(); p != nil {
				if p == hijack {
					return
				}
				panic(p)
			}
		}()
		handler.ServeHTTP(hijack, r)
	})
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	code := parseCodeRequestParam(r.FormValue("t"))
	if code == "" {
		code = exampleCode
	}

	result, err := discoverCode(code)
	errStr := ""
	if err != nil {
		errStr = err.Error()
	}

	templates["index"].ExecuteTemplate(
		w, "base", &struct {
			Code   string
			Result string
			Error  string
		}{code, result, errStr},
	) // todo: check error
}

func parseCodeRequestParam(param string) string {
	param = strings.TrimSpace(param)
	bytes, err := base64.URLEncoding.DecodeString(param)
	if err != nil {
		return ""
	}
	return string(bytes)
}
