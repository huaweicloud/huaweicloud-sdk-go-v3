package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataspacesRequest Request Object
type ListDataspacesRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// dataspace name
	DataspaceName *string `json:"dataspace_name,omitempty"`

	// 数据空间ID
	DataspaceId *string `json:"dataspace_id,omitempty"`

	// 第几页
	Offset *int64 `json:"offset,omitempty"`

	// 每页数量
	Limit *int64 `json:"limit,omitempty"`

	// 排序顺序
	SortDir *string `json:"sort_dir,omitempty"`

	// 排序字段
	SortKey *string `json:"sort_key,omitempty"`
}

func (o ListDataspacesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataspacesRequest struct{}"
	}

	return strings.Join([]string{"ListDataspacesRequest", string(data)}, " ")
}
