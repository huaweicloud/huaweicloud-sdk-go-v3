/*
 * NAT
 *
 * Open Api of Public Nat.
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchCreateDnatRulesRequest struct {
	Body *BatchCreateDnatRulesRequestBody `json:"body,omitempty"`
}

func (o BatchCreateDnatRulesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchCreateDnatRulesRequest", string(data)}, " ")
}
