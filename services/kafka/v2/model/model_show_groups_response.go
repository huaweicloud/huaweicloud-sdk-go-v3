/*
 * Kafka
 *
 * Kafka Document API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowGroupsResponse struct {
	Group          *ShowGroupsRespGroup `json:"group,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowGroupsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowGroupsResponse", string(data)}, " ")
}
