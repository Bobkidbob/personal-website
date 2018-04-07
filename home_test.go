package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type homeTest struct {
	*defaultRouteTest
}

func (home *homeTest) TestGet(t *testing.T) {
	testHome(t, http.MethodGet, home.Path)
}

func (home *homeTest) TestHead(t *testing.T) {
	testHome(t, http.MethodHead, home.Path)
}

func testHome(t *testing.T, method, path string) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(method, "https://www.bobkidbob.com"+path, nil)
	handler(w, r)
	res := w.Result()
	if e, g := "303 See Other", res.Status; e != g {
		t.Errorf("Expected HTTP status \"%v\", but got \"%v\"", e, g)
	}
	l, err := res.Location()
	if err != nil {
		t.Error(err)
	}
	if e, g := "https://www.linkedin.com/in/bobkidbob/", l.String(); e != g {
		t.Errorf("Expected location header \"%v\", but got \"%v\"", e, g)
	}
}
