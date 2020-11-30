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
type UpdateNickNameV4Response struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateNickNameV4Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateNickNameV4Response", string(data)}, " ")
}
