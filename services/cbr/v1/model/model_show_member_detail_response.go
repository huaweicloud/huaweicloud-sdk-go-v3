package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowMemberDetailResponse struct {
	Member         *Member `json:"member,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowMemberDetailResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowMemberDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowMemberDetailResponse", string(data)}, " ")
}
