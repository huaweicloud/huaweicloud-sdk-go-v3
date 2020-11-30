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
type BatchDeleteMembersV4Response struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteMembersV4Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchDeleteMembersV4Response", string(data)}, " ")
}
