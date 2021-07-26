package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowLtsLogDetailsRequest struct {
	// 函数的URN（Uniform Resource Name），唯一标识函数。

	FunctionUrn string `json:"function_urn"`
}

func (o ShowLtsLogDetailsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowLtsLogDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowLtsLogDetailsRequest", string(data)}, " ")
}
