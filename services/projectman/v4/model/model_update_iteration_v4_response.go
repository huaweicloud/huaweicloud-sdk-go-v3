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
type UpdateIterationV4Response struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateIterationV4Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateIterationV4Response", string(data)}, " ")
}
