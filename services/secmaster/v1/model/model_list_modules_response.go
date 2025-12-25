package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListModulesResponse Response Object
type ListModulesResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 数据列表
	Data *[]ModuleDetailInfo `json:"data,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	// 页码
	Offset *int32 `json:"offset,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	// 数量
	Limit *int32 `json:"limit,omitempty"`

	// 是否响应成功
	Success *bool `json:"success,omitempty"`

	// 总数
	Total *int32 `json:"total,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListModulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListModulesResponse struct{}"
	}

	return strings.Join([]string{"ListModulesResponse", string(data)}, " ")
}
