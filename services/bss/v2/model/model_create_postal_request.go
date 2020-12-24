/*
 * BSS
 *
 * Business Support System API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreatePostalRequest struct {
	XLanguage *string       `json:"X-Language,omitempty"`
	Body      *AddPostalReq `json:"body,omitempty"`
}

func (o CreatePostalRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreatePostalRequest", string(data)}, " ")
}
