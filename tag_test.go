// Copyright (c) 2022, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package flagtag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTag(t *testing.T) {
	tests := map[string]struct {
		input   string
		want    Name
		wantErr error
	}{
		"": {},
		"short": {
			input: "h",
			want:  Name{Short: "h"},
		},
		"long": {
			input: "help",
			want:  Name{Long: "help"},
		},
		"short,long": {
			input: "h,help",
			want:  Name{Short: "h", Long: "help"},
		},
		"long,short": {
			input: "help,h",
			want:  Name{Short: "h", Long: "help"},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			have, _ := ParseTag(tc.input)
			assert.Equal(t, tc.want, have)
		})
	}
}
