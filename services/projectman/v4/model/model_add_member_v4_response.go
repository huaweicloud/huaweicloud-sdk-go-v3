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
type AddMemberV4Response struct {
}

func (o AddMemberV4Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"AddMemberV4Response", string(data)}, " ")
}
