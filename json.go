// Copyright 2022 youwkey. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package conv

import (
	"bytes"
	"encoding/json"

	"github.com/youwkey/alfred-go"
)

// JSONMinify minify json string and return it set to alfred.Item.
func JSONMinify(src string) *alfred.Item {
	var buf bytes.Buffer
	if err := json.Compact(&buf, []byte(src)); err != nil {
		return alfred.NewInvalidItem("JSON Minify Error").Subtitle(err.Error())
	}

	str := buf.String()

	return alfred.NewItem(str).Text(str).Arg(str)
}

// JSONPretty pretty json string and return it set to alfred.Item.
func JSONPretty(src string) *alfred.Item {
	var buf bytes.Buffer
	if err := json.Indent(&buf, []byte(src), "", "  "); err != nil {
		return alfred.NewInvalidItem("JSON Pretty Error").Subtitle(err.Error())
	}

	str := buf.String()

	return alfred.NewItem(str).Text(str).Arg(str)
}
