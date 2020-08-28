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
type CreateProjectV4Request struct {
	Body *CreateProjectV4RequestBody `json:"body,omitempty"`
}

func (o CreateProjectV4Request) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateProjectV4Request", string(data)}, " ")
}
