// Copyright (c) 2022, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package flagtag

import (
	"strings"

	"github.com/go-pogo/errors"
)

const (
	ErrMultipleShortNames errors.Msg = "found multiple short names"
	ErrMultipleLongNames  errors.Msg = "found multiple long names"
)

type Name struct {
	Long  string
	Short string
}

// ParseTag parses tag to a Name containing at least a long or short name.
func ParseTag(tag string) (Name, error) {
	if tag == "" {
		return Name{}, nil
	}

	split := strings.Split(tag, ",")
	splitN := len(split)

	if splitN == 1 {
		var name Name
		name.set(split[0])
		return name, nil
	}

	var name Name
	nt := name.set(split[0]) + name.set(split[1])
	if nt == short+short {
		return Name{}, errors.Wrap(ErrMultipleShortNames, &TagError{Tag: tag})
	} else if nt == long+long {
		return Name{}, errors.Wrap(ErrMultipleLongNames, &TagError{Tag: tag})
	}

	return name, nil
}

type TagError struct {
	Tag string
}

func (e *TagError) Error() string { return "invalid tag '" + e.Tag + "'" }

// nameType is used to check if a provided flag name is short or long.
// It is important that the const value of long is not equal to short + short
// and, likewise short is not equal to long + long.
type nameType int8

const (
	none  nameType = 0
	short nameType = 1
	long  nameType = 3
)

func (fn *Name) set(str string) nameType {
	if str == "" {
		return none
	}

	str = strings.TrimSpace(str)
	if str == "" {
		return none
	}
	if len(str) == 1 {
		fn.Short = str
		return short
	}

	fn.Long = str
	return long
}
