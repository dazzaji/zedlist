// Copyright 2015 Geofrey Ernest <geofreyernest@live.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package auth

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gernest/zedlist/modules/forms"
	"github.com/gernest/zedlist/modules/query"
	"github.com/gernest/zedlist/modules/settings"

	"github.com/labstack/echo"
)

func TestNormal(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	e := echo.New()
	ctx := echo.NewContext(req, echo.NewResponse(w), e)
	err = Normal()(ctx)
	if err != nil {
		t.Errorf("expected nil got %v", err)
	}
}

func TestMust(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	e := echo.New()
	ctx := echo.NewContext(req, echo.NewResponse(w), e)
	err = Must()(ctx)
	if err == nil {
		t.Error("expected an error got nil")
	}
	if !strings.Contains(err.Error(), "Unauthorized") {
		t.Errorf("expected Unauthorized got %v", err)
	}

	regForm := forms.Register{
		FirstName:       "root",
		LastName:        "roote",
		MiddleName:      "root",
		Email:           "auth@home.com",
		Password:        "superroot",
		ConfirmPassword: "superroot",
		Gender:          1,
		BirthDay:        time.Now(),
	}
	usr, err := query.CreateNewUser(regForm)
	if err != nil {
		t.Errorf("creating new user %v", err)
	}

	defer query.Delete(usr)

	ctx.Set("IsLoged", true)
	err = Must()(ctx)
	if err != nil {
		t.Errorf("expected nil got %v", err)
	}
	nctx := echo.NewContext(req, echo.NewResponse(w), e)
	ss, _ := store.Get(nctx.Request(), settings.App.Session.Name)
	ss.Values["userID"] = usr.ID
	err = ss.Save(nctx.Request(), nctx.Response())
	if err != nil {
		t.Error(err)
	}
	err = Must()(nctx)
	if err != nil {
		t.Errorf("expected nil got %v", err)
	}
}
