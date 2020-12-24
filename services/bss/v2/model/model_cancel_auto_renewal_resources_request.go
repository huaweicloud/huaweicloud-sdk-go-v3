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
type CancelAutoRenewalResourcesRequest struct {
	ResourceId string `json:"resource_id"`
}

func (o CancelAutoRenewalResourcesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CancelAutoRenewalResourcesRequest", string(data)}, " ")
}
