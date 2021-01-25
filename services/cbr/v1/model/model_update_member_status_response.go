package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateMemberStatusResponse struct {
	Member         *Member `json:"member,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateMemberStatusResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateMemberStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdateMemberStatusResponse", string(data)}, " ")
}
