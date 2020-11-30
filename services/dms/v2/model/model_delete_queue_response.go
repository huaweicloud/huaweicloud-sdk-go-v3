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
type DeleteQueueResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteQueueResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteQueueResponse", string(data)}, " ")
}
