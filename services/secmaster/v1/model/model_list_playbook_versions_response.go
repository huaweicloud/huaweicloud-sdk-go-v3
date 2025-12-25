package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPlaybookVersionsResponse Response Object
type ListPlaybookVersionsResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	// 分页查询数据大小
	Size *int32 `json:"size,omitempty"`

	// 当前页码
	Page *int32 `json:"page,omitempty"`

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 剧本版本列表信息
	Data *[]PlaybookVersionListEntity `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPlaybookVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPlaybookVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListPlaybookVersionsResponse", string(data)}, " ")
}
