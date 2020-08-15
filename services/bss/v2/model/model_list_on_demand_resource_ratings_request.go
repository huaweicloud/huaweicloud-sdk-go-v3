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
type ListOnDemandResourceRatingsRequest struct {
	Body *RateOnDemandReq `json:"body,omitempty"`
}

func (o ListOnDemandResourceRatingsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListOnDemandResourceRatingsRequest", string(data)}, " ")
}
