package app

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/gophergala/golang-sizeof.tips/internal/parser"
)

const exampleCode = `
struct {
	a string
	b bool
	c string
}
`

func discoverHandler(w http.ResponseWriter, r *http.Request) {
	code := parseCodeRequestParam(r.FormValue("t"))
	if code == "" {
		code = exampleCode
	}

	result, err := parser.ParseCode(code)
	errStr := ""
	if err != nil {
		errStr = err.Error()
	}

	templates["index"].ExecuteTemplate(
		w, "base", &struct {
			Code   string
			Result string
			Error  string
		}{code, fmt.Sprintf("%+v", result), errStr},
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
