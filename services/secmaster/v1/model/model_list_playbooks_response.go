package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPlaybooksResponse Response Object
type ListPlaybooksResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 响应消息信息
	Message *string `json:"message,omitempty"`

	// 总条数
	Total *int32 `json:"total,omitempty"`

	// 分页查询数据大小
	Size *int32 `json:"size,omitempty"`

	// 当前页码
	Page *int32 `json:"page,omitempty"`

	// 剧本列表信息
	Data *[]PlaybookInfo `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPlaybooksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPlaybooksResponse struct{}"
	}

	return strings.Join([]string{"ListPlaybooksResponse", string(data)}, " ")
}
