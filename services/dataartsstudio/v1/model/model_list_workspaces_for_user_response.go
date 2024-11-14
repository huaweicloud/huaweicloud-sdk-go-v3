package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkspacesForUserResponse Response Object
type ListWorkspacesForUserResponse struct {

	// 当前工作空间用户记录数
	Count *int32 `json:"count,omitempty"`

	// 查询结果总页数
	TotalPage *int32 `json:"total_page,omitempty"`

	// 工作空间列表
	Data           *[]Workspacebody `json:"data,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListWorkspacesForUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkspacesForUserResponse struct{}"
	}

	return strings.Join([]string{"ListWorkspacesForUserResponse", string(data)}, " ")
}
