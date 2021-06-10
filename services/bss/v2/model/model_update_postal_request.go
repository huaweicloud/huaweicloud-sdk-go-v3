package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdatePostalRequest struct {
	// |参数名称：语言| |参数的约束及描述：中文：zh_CN 英文：en_US缺省为zh_CN|

	XLanguage *string `json:"X-Language,omitempty"`

	Body *UpdatePostalReq `json:"body,omitempty"`
}

func (o UpdatePostalRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePostalRequest struct{}"
	}

	return strings.Join([]string{"UpdatePostalRequest", string(data)}, " ")
}
