package actions

import (
	"fmt"
	"github.com/pinchtab/pinchtab/internal/cli"
	"github.com/pinchtab/pinchtab/internal/cli/apiclient"
	"net/http"
	"net/url"
	"os"
	"time"
)

func Screenshot(client *http.Client, base, token string, args []string) {
	params := url.Values{}
	params.Set("raw", "true")
	outFile := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-o", "--output":
			if i+1 < len(args) {
				i++
				outFile = args[i]
			}
		case "--quality", "-q":
			if i+1 < len(args) {
				i++
				params.Set("quality", args[i])
			}
		case "--tab":
			if i+1 < len(args) {
				i++
				params.Set("tabId", args[i])
			}
		}
	}

	if outFile == "" {
		outFile = fmt.Sprintf("screenshot-%s.jpg", time.Now().Format("20060102-150405"))
	}

	data := apiclient.DoGetRaw(client, base, token, "/screenshot", params)
	if data == nil {
		return
	}
	if err := os.WriteFile(outFile, data, 0600); err != nil {
		cli.Fatal("Write failed: %v", err)
	}
	fmt.Println(cli.StyleStdout(cli.SuccessStyle, fmt.Sprintf("Saved %s (%d bytes)", outFile, len(data))))
}
