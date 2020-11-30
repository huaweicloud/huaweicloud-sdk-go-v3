/*
 * ProjectMan
 *
 * devcloud projectman api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchDeleteIterationsV4Response struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteIterationsV4Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchDeleteIterationsV4Response", string(data)}, " ")
}
