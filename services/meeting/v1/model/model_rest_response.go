package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Rest请求的响应结果body
type RestResponse struct {
	// 结果码

	ReturnCode int32 `json:"returnCode"`
	// 结果描述

	ReturnDesc *string `json:"returnDesc,omitempty"`
}

func (o RestResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestResponse struct{}"
	}

	return strings.Join([]string{"RestResponse", string(data)}, " ")
}
