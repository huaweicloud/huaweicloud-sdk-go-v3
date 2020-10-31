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

// Request Object
type CheckProjectNameV4Request struct {
	Body *CheckProjectNameRequestV4 `json:"body,omitempty"`
}

func (o CheckProjectNameV4Request) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CheckProjectNameV4Request", string(data)}, " ")
}
