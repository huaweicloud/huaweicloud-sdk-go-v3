package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProtectionConfigurationResponseItem struct {

	// 资源ID
	Ids []string `json:"ids"`

	// 状态码
	Code int32 `json:"code"`

	// 错误编码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`

	// bcc向其他云服务请求的requestId
	RequestId *string `json:"request_id,omitempty"`

	// 发生错误的阶段
	ErrorStage *string `json:"error_stage,omitempty"`
}

func (o ProtectionConfigurationResponseItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectionConfigurationResponseItem struct{}"
	}

	return strings.Join([]string{"ProtectionConfigurationResponseItem", string(data)}, " ")
}
