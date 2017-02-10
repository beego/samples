// Copyright 2017 Beego Samples authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package controllers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/astaxie/beego"
	"golang.org/x/net/webdav"
)

var (
	path   = "_testdata"
	prefix = "/"
)

// WebDav WebDav struct
type WebDav struct {
	fs          webdav.FileSystem
	ls          webdav.LockSystem
	HandlerFunc http.HandlerFunc
}

// WebDAVController handles WebDAV requests.
type WebDAVController struct {
	beego.Controller
}

func NewWebDav() *WebDav {
	return &WebDav{}
}

func (wd *WebDav) mount(path string) error {
	if s, err := filepath.Abs(path); err == nil {
		path = s
	}
	wd.fs = webdav.Dir(path)
	wd.ls = webdav.NewMemLS()
	return nil
}

// Main All method handles all requests for WebDAVController.
func (c *WebDAVController) Main() {
	wd := NewWebDav()
	wd.mount(path)

	os.Mkdir(path, os.ModeDir)

	h := &webdav.Handler{
		FileSystem: wd.fs,
		LockSystem: wd.ls,
		Prefix:     prefix,
	}

	h.ServeHTTP(c.Ctx.ResponseWriter, c.Ctx.Request)
}
