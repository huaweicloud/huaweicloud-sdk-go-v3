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
type DeleteIterationV4Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteIterationV4Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteIterationV4Response", string(data)}, " ")
}
