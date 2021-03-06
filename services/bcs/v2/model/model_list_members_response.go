package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListMembersResponse struct {
	// 联盟成员列表

	Members        *[]Member `json:"members,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListMembersResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListMembersResponse struct{}"
	}

	return strings.Join([]string{"ListMembersResponse", string(data)}, " ")
}
