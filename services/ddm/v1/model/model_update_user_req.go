package model

import (
	"encoding/json"

	"strings"
)

// This is a auto update request Object
type UpdateUserReq struct {
	User *UpdateUserDetailReq `json:"user"`
}

func (o UpdateUserReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateUserReq struct{}"
	}

	return strings.Join([]string{"UpdateUserReq", string(data)}, " ")
}
