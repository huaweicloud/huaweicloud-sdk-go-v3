package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LayoutCommonRsp struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	// 请求结果
	Success *bool `json:"success,omitempty"`
}

func (o LayoutCommonRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LayoutCommonRsp struct{}"
	}

	return strings.Join([]string{"LayoutCommonRsp", string(data)}, " ")
}
