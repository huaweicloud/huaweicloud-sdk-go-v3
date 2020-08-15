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
type ListCustomerselfResourceRecordDetailsRequest struct {
	Body *QueryResRecordsDetailReq `json:"body,omitempty"`
}

func (o ListCustomerselfResourceRecordDetailsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListCustomerselfResourceRecordDetailsRequest", string(data)}, " ")
}
