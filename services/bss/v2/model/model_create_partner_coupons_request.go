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
type CreatePartnerCouponsRequest struct {
	Body *CreatePartnerCouponsReq `json:"body,omitempty"`
}

func (o CreatePartnerCouponsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreatePartnerCouponsRequest", string(data)}, " ")
}
