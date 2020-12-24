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
type ListSubCustomersRequest struct {
	Body *QuerySubCustomerListReq `json:"body,omitempty"`
}

func (o ListSubCustomersRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListSubCustomersRequest", string(data)}, " ")
}
