package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkspaceusersResponse Response Object
type ListWorkspaceusersResponse struct {

	// 当前工作空间用户记录数
	Count *int32 `json:"count,omitempty"`

	// 查询结果总页数
	TotalPage *int32 `json:"total_page,omitempty"`

	// 工作空间用户列表
	Data           *[]ApigWorkspaceUserbody `json:"data,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListWorkspaceusersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkspaceusersResponse struct{}"
	}

	return strings.Join([]string{"ListWorkspaceusersResponse", string(data)}, " ")
}
