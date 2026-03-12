package actions

import (
	"github.com/pinchtab/pinchtab/internal/cli"
	"github.com/pinchtab/pinchtab/internal/cli/apiclient"
	"net/http"
	"strings"
)

func Evaluate(client *http.Client, base, token string, args []string) {
	if len(args) < 1 {
		cli.Fatal("Usage: pinchtab eval <expression>")
	}
	expr := strings.Join(args, " ")
	apiclient.DoPost(client, base, token, "/evaluate", map[string]any{
		"expression": expr,
	})
}
