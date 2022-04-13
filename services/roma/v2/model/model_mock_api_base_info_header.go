package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MockApiBaseInfoHeader struct {
	// mock后端自定义响应头header key

	Key *string `json:"key,omitempty"`
	// mock后端自定义响应头header value

	Value *string `json:"value,omitempty"`
	// mock后端自定义响应头header remark

	Remark *string `json:"remark,omitempty"`
}

func (o MockApiBaseInfoHeader) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MockApiBaseInfoHeader struct{}"
	}

	return strings.Join([]string{"MockApiBaseInfoHeader", string(data)}, " ")
}
