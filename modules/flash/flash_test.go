// Copyright 2015 Geofrey Ernest <geofreyernest@live.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package flash

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
)

func TestFlash(t *testing.T) {

	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Error(err)
	}
	client := &http.Client{Jar: jar}
	e := echo.New()
	e.Get("/", func(ctx *echo.Context) error {
		fls := New()
		fls.Success("Success")
		fls.Err("Err")
		fls.Warn("Warn")
		fls.Save(ctx)
		return nil
	})

	var result Flashes

	e.Get("/flash", func(ctx *echo.Context) error {
		result = GetFlashes(ctx)
		return nil
	})
	ts := httptest.NewServer(e)
	defer ts.Close()

	_, err = client.Get(fmt.Sprintf("%s/", ts.URL))
	if err != nil {
		t.Error(err)
	}
	_, err = client.Get(fmt.Sprintf("%s/flash", ts.URL))
	if err != nil {
		t.Error(err)
	}
	if result[0].Message != "Success" {
		t.Errorf("expected success got %s", result[0].Message)
	}
	if len(result) != 3 {
		t.Errorf("expected 3 got %d", len(result))
	}
}
