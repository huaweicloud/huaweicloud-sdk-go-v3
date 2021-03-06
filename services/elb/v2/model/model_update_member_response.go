package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateMemberResponse struct {
	Member         *MemberResp `json:"member,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o UpdateMemberResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateMemberResponse struct{}"
	}

	return strings.Join([]string{"UpdateMemberResponse", string(data)}, " ")
}
