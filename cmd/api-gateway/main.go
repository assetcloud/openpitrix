// Copyright 2017 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

// openpitrix api gateway
package main

import (
	"openpitrix.io/openpitrix/pkg/cmd/api"
)

func main() {
	api.Serve()
}
