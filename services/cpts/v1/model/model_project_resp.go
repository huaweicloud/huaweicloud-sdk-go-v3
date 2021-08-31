package model

import (
	"encoding/json"

	"strings"
)

type ProjectResp struct {
	// 状态码

	Code *string `json:"code,omitempty"`
	// 描述

	Message *string `json:"message,omitempty"`
}

func (o ProjectResp) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ProjectResp struct{}"
	}

	return strings.Join([]string{"ProjectResp", string(data)}, " ")
}
