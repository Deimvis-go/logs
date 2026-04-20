package logs

import (
	"fmt"
	"net/http"

	"github.com/Deimvis/go-ext/go1.25/xhttp"
	"go.uber.org/zap"
)

func NewZapRoundTripWrap(z *zap.SugaredLogger) xhttp.RoundTripWrapFn {
	lg := ZapAsKVCtxLogger(z)
	return func(rtFn xhttp.RoundTripFn) xhttp.RoundTripFn {
		return func(req *http.Request) (*http.Response, error) {
			ctx := req.Context()
			reqStr := formatRequest(req)
			lg.Debug(ctx, "Subrequest start", "request", reqStr)
			resp, err := rtFn(req)
			if err != nil {
				lg.Debug(ctx, "Subrequest fail", "request", reqStr, "error", err)
			} else {
				lg.Debug(ctx, "Subrequest finish", "response", formatResponse(resp), "request", reqStr)
			}
			return resp, err
		}
	}
}

func formatRequest(req *http.Request) string {
	return fmt.Sprintf("%s %s", req.Method, req.URL.String())
}

func formatResponse(resp *http.Response) string {
	return fmt.Sprintf("%d", resp.StatusCode)
}
