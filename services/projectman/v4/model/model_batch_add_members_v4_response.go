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
type BatchAddMembersV4Response struct {
}

func (o BatchAddMembersV4Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchAddMembersV4Response", string(data)}, " ")
}
