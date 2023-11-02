package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPlaybookActionsRequest Request Object
type ListPlaybookActionsRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 剧本版本ID
	VersionId string `json:"version_id"`

	// 分页查询参数，用于指定一次查询最多的结果数，从1开始
	Limit int32 `json:"limit"`

	// 分页查询参数。用于指定查询结果的起始位置，从0开始
	Offset int32 `json:"offset"`
}

func (o ListPlaybookActionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPlaybookActionsRequest struct{}"
	}

	return strings.Join([]string{"ListPlaybookActionsRequest", string(data)}, " ")
}
