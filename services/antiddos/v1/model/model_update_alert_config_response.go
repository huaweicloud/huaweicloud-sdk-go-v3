package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateAlertConfigResponse struct {
	// 内部错误码

	ErrorCode *string `json:"error_code,omitempty"`
	// 内部错误描述

	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAlertConfigResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateAlertConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateAlertConfigResponse", string(data)}, " ")
}
