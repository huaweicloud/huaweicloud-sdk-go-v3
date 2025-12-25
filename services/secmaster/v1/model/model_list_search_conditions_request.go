package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSearchConditionsRequest Request Object
type ListSearchConditionsRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 管道ID
	PipeId string `json:"pipe_id"`

	// 第几页
	Offset *int64 `json:"offset,omitempty"`

	// 每页数量
	Limit *int64 `json:"limit,omitempty"`

	// 排序字段
	SortKey *string `json:"sort_key,omitempty"`

	// 排序顺序
	SortDir *string `json:"sort_dir,omitempty"`
}

func (o ListSearchConditionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSearchConditionsRequest struct{}"
	}

	return strings.Join([]string{"ListSearchConditionsRequest", string(data)}, " ")
}
