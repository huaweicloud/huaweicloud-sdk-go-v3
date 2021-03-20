package model

import (
	"encoding/json"

	"strings"
)

type AuditReport struct {
	// 疑似黑边位置信息，参数格式：top:bottom:left:right

	BlackPosition *string `json:"black_position,omitempty"`
}

func (o AuditReport) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AuditReport struct{}"
	}

	return strings.Join([]string{"AuditReport", string(data)}, " ")
}
