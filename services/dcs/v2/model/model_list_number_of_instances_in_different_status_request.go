/*
 * DCS
 *
 * DCS V2版本API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListNumberOfInstancesInDifferentStatusRequest struct {
	IncludeFailure *string `json:"include_failure,omitempty"`
}

func (o ListNumberOfInstancesInDifferentStatusRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListNumberOfInstancesInDifferentStatusRequest", string(data)}, " ")
}
