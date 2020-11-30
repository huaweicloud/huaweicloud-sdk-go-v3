/*
 * DMS
 *
 * DMS Document API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteSpecifiedConsumerGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSpecifiedConsumerGroupResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteSpecifiedConsumerGroupResponse", string(data)}, " ")
}
