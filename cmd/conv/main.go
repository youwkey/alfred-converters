// Copyright 2022 youwkey. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"flag"

	"github.com/youwkey/alfred-go"
	conv "github.com/youwkey/alfred-mutual-converter"
)

func main() {
	flag.Parse()
	action, arg := flag.Arg(0), flag.Arg(1)
	sf := alfred.NewScriptFilter()

	if action == "" || arg == "" {
		sf.Items().Append(alfred.NewInvalidItem(action + " '...'"))
		_ = sf.Output()

		return
	}

	switch action {
	case "json/pretty":
		sf.Items().Append(conv.JSONPretty(arg))
	case "json/minify":
		sf.Items().Append(conv.JSONMinify(arg))
	}

	_ = sf.OutputIndent("", "  ")
}
