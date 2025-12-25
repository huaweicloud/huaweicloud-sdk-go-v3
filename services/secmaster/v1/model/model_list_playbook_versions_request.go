package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPlaybookVersionsRequest Request Object
type ListPlaybookVersionsRequest struct {

	// **参数解释：** 内容类型 - application/json    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json  **默认取值：** application/json
	ContentType string `json:"content-type"`

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 剧本ID
	PlaybookId string `json:"playbook_id"`

	// 剧本版本状态，编辑中：EDITING 审核中：APPROVING 不通过：UNPASSED 已发布：PUBLISHED
	Status *string `json:"status,omitempty"`

	// 启用/禁用
	Enabled *int32 `json:"enabled,omitempty"`

	// 版本类型， 草稿版本：0 正式版本：1
	VersionType *int32 `json:"version_type,omitempty"`

	// 分页查询参数。用于指定查询结果的起始位置，从0开始
	Offset *int32 `json:"offset,omitempty"`

	// 分页查询参数，用于指定一次查询最多的结果数，从1开始
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListPlaybookVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPlaybookVersionsRequest struct{}"
	}

	return strings.Join([]string{"ListPlaybookVersionsRequest", string(data)}, " ")
}
