// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package staticheaders

import (
	"github.com/google/go-safeweb/safehttp"
)

// Plugin claims and sets static headers on responses.
type Plugin struct{}

// Before claims and sets the following headers:
//  - X-Content-Type-Options: nosniff
//  - X-XSS-Protection: 0
func (Plugin) Before(w safehttp.ResponseWriter, r *safehttp.IncomingRequest) safehttp.Result {
	h := w.Header()
	setXCTO, err := h.Claim("X-Content-Type-Options")
	if err != nil {
		return w.ServerError(safehttp.StatusInternalServerError)
	}
	setXCTO([]string{"nosniff"})

	setXXP, err := h.Claim("X-XSS-Protection")
	if err != nil {
		return w.ServerError(safehttp.StatusInternalServerError)
	}
	setXXP([]string{"0"})
	return safehttp.Result{}
}
