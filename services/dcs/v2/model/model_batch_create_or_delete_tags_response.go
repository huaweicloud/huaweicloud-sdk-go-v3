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

// Response Object
type BatchCreateOrDeleteTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateOrDeleteTagsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchCreateOrDeleteTagsResponse", string(data)}, " ")
}
