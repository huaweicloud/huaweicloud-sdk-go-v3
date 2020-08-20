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
type ListCustomerOnDemandResourcesRequest struct {
	Body *QueryCustomerOnDemandResourcesReq `json:"body,omitempty"`
}

func (o ListCustomerOnDemandResourcesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListCustomerOnDemandResourcesRequest", string(data)}, " ")
}
