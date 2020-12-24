/*
 * SMN
 *
 * SMN Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListApplicationsRequest struct {
	Offset   *int32  `json:"offset,omitempty"`
	Limit    *int32  `json:"limit,omitempty"`
	Name     *string `json:"name,omitempty"`
	Platform *string `json:"platform,omitempty"`
}

func (o ListApplicationsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListApplicationsRequest", string(data)}, " ")
}
