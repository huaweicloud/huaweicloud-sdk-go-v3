package model

import (
	"encoding/json"

	"strings"
)

type WarningList struct {
	// 警告ID。

	WarningCode *int32 `json:"warningCode,omitempty"`
	// 告警消息。

	WarningMsg *string `json:"warningMsg,omitempty"`
}

func (o WarningList) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "WarningList struct{}"
	}

	return strings.Join([]string{"WarningList", string(data)}, " ")
}
