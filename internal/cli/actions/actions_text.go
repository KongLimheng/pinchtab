package actions

import (
	"github.com/pinchtab/pinchtab/internal/cli/apiclient"
	"net/http"
	"net/url"
)

func Text(client *http.Client, base, token string, args []string) {
	params := url.Values{}
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--raw":
			params.Set("mode", "raw")
		case "--tab":
			if i+1 < len(args) {
				i++
				params.Set("tabId", args[i])
			}
		}
	}
	apiclient.DoGet(client, base, token, "/text", params)
}
