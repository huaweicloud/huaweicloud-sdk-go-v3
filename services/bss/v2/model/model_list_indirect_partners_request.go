/*
 * Bss
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
type ListIndirectPartnersRequest struct {
	Body *QueryIndirectPartnersReq `json:"body,omitempty"`
}

func (o ListIndirectPartnersRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListIndirectPartnersRequest", string(data)}, " ")
}
