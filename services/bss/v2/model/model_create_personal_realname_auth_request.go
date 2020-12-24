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
type CreatePersonalRealnameAuthRequest struct {
	Body *ApplyIndividualRealnameAuthsReq `json:"body,omitempty"`
}

func (o CreatePersonalRealnameAuthRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreatePersonalRealnameAuthRequest", string(data)}, " ")
}
