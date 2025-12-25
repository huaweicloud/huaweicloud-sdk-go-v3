package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateModuleResponse Response Object
type UpdateModuleResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	// 请求结果
	Success *bool `json:"success,omitempty"`

	Data *ModuleDetailInfo `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateModuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateModuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateModuleResponse", string(data)}, " ")
}
