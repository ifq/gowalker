// Copyright 2011 Gary Burd
// Copyright 2013-2014 Unknown
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

package doc

import (
	"errors"

	"github.com/Unknwon/com"
)

var (
	errNotModified   = errors.New("Package not modified")
	errNoMatch       = errors.New("no match")
	errUpdateTimeout = errors.New("update timeout")
)

func isNotFound(err error) bool {
	_, ok := err.(com.NotFoundError)
	return ok
}
