// httpb パッケージは HTTP client のユーティリティを提供します.
package httpb

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// StatusError は HTTP 通信でエラー扱いになるステータスを表します.
type StatusError struct {
	Method     string
	URL        string
	StatusCode int
}

// ErrStatus は StatusError を作成します.
func ErrStatus(resp *http.Response) StatusError {
	return StatusError{
		Method:     resp.Request.Method,
		URL:        resp.Request.URL.String(),
		StatusCode: resp.StatusCode,
	}
}

func (e StatusError) Error() string {
	method := bytes.ToLower([]byte(e.Method))
	method[0] -= 'a' - 'A'

	return fmt.Sprintf(
		"%s %q: %s - %d",
		method,
		e.URL,
		strings.ToLower(http.StatusText(e.StatusCode)),
		e.StatusCode,
	)
}

// GetLastModified は http.Response から更新日時を取得します.
//
// 取得した更新日時は time.Local のタイムゾーンに変換しています.
func GetLastModified(resp *http.Response) time.Time {
	tm, err := time.Parse(time.RFC1123, resp.Header.Get("Last-Modified"))
	if err != nil {
		tm = time.Now()
	}
	return tm.Local()
}
