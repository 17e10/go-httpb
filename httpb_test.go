package httpb

import (
	"net/http"
	"testing"
	"time"
)

func TestErrStatus(t *testing.T) {
	resp, err := http.Get("https://go.dev/")
	if err != nil {
		t.Fatal(err)
	}
	err = ErrStatus(resp)
	got := err.Error()
	want := `Get "https://go.dev/": ok - 200`
	if got != want {
		t.Errorf("error = %q, want %q", got, want)
	}
}

func TestLastModified(t *testing.T) {
	gmt, err := time.LoadLocation("GMT")
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.Get("https://go.dev/")
	if err != nil {
		t.Fatal(err)
	}

	// Last-Modified に 5 分前を設定してテストする
	want := time.Now().Add(-300 * time.Second).In(gmt)
	resp.Header.Set("Last-Modified", want.Format(time.RFC1123))
	got := GetLastModified(resp)
	if got.Unix() != want.Unix() {
		t.Errorf("Last-Modified = %s, want %s", got.Format(time.DateTime), want.Local().Format(time.DateTime))
	}
}
