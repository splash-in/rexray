// +build !go1.7
// +build !rexray_build_type_agent
// +build !rexray_build_type_controller

package scripts

import (
	"net/http"

	apitypes "github.com/codedellemc/rexray/libstorage/api/types"
	"golang.org/x/net/context/ctxhttp"
)

func doRequest(
	ctx apitypes.Context, req *http.Request) (*http.Response, error) {
	return doRequestWithClient(ctx, http.DefaultClient, req)
}

func doRequestWithClient(
	ctx apitypes.Context,
	client *http.Client,
	req *http.Request) (*http.Response, error) {
	return ctxhttp.Do(ctx, client, req)
}
