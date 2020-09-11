/*
 * kps
 *
 * kps v3 版本API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type AssociateKeypairRequest struct {
	Body *AssociateKeypairRequestBody `json:"body,omitempty"`
}

func (o AssociateKeypairRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"AssociateKeypairRequest", string(data)}, " ")
}
