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
type UpdateMembesRoleV4Response struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateMembesRoleV4Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateMembesRoleV4Response", string(data)}, " ")
}
