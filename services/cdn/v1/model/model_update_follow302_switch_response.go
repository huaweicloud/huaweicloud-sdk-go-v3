package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateFollow302SwitchResponse struct {
	FollowStatus   *Follow302StatusBody `json:"follow_status,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o UpdateFollow302SwitchResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateFollow302SwitchResponse struct{}"
	}

	return strings.Join([]string{"UpdateFollow302SwitchResponse", string(data)}, " ")
}
