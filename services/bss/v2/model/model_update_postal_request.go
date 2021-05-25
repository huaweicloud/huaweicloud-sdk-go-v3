package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdatePostalRequest struct {
	// 语言。 中文：zh_CN 缺省为zh_CN。

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
