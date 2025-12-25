package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowModuleResponse Response Object
type ShowModuleResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	Data *ModuleDetailInfo `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowModuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowModuleResponse struct{}"
	}

	return strings.Join([]string{"ShowModuleResponse", string(data)}, " ")
}
