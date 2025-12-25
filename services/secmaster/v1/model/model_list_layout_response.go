package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLayoutResponse Response Object
type ListLayoutResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 数据列表
	Data *[]LayoutDetailInfo `json:"data,omitempty"`

	// 响应信息
	Message *string `json:"message,omitempty"`

	// 分页
	Offset *int32 `json:"offset,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	// 每页个数
	Limit *int32 `json:"limit,omitempty"`

	// 是否响应成功
	Success *bool `json:"success,omitempty"`

	// 总数
	Total *int32 `json:"total,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListLayoutResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLayoutResponse struct{}"
	}

	return strings.Join([]string{"ListLayoutResponse", string(data)}, " ")
}
