package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPlaybookActionsResponse Response Object
type ListPlaybookActionsResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 分页大小
	Size *int32 `json:"size,omitempty"`

	// 当前页数
	Page *int32 `json:"page,omitempty"`

	// 剧本动作列表信息
	Data *[]ActionInfo `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPlaybookActionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPlaybookActionsResponse struct{}"
	}

	return strings.Join([]string{"ListPlaybookActionsResponse", string(data)}, " ")
}
