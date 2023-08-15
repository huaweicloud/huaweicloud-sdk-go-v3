package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkspaceusersRequest Request Object
type ListWorkspaceusersRequest struct {

	// 工作空间id
	WorkspaceId string `json:"workspace_id"`

	// 数据条数限制
	Limit *string `json:"limit,omitempty"`

	// 偏移量
	Offset *string `json:"offset,omitempty"`
}

func (o ListWorkspaceusersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkspaceusersRequest struct{}"
	}

	return strings.Join([]string{"ListWorkspaceusersRequest", string(data)}, " ")
}
