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
type BatchCreateOrDeleteQueueTagResponse struct {
}

func (o BatchCreateOrDeleteQueueTagResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchCreateOrDeleteQueueTagResponse", string(data)}, " ")
}
