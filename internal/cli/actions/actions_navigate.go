package actions

import (
	"github.com/pinchtab/pinchtab/internal/cli"
	"github.com/pinchtab/pinchtab/internal/cli/apiclient"
	"net/http"
)

func Navigate(client *http.Client, base, token string, args []string) {
	if len(args) < 1 {
		cli.Fatal("Usage: pinchtab nav <url> [--new-tab] [--block-images] [--block-ads]")
	}
	body := map[string]any{"url": args[0]}
	for _, a := range args[1:] {
		switch a {
		case "--new-tab":
			body["newTab"] = true
		case "--block-images":
			body["blockImages"] = true
		case "--block-ads":
			body["blockAds"] = true
		}
	}
	result := apiclient.DoPost(client, base, token, "/navigate", body)
	apiclient.SuggestNextAction("navigate", result)
}
