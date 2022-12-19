package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPlaybookVersionsRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// ID of workspace
	WorkspaceId string `json:"workspace_id"`

	// ID of playbook
	PlaybookId string `json:"playbook_id"`

	// 剧本版本状态，编辑中：EDITING  审核中：APPROVING  不通过：UNPASSED 已发布：PUBLISHED
	Status *string `json:"status,omitempty"`

	// 启用/禁用
	Enabled *int32 `json:"enabled,omitempty"`

	// 版本类型， 草稿版本：0  正式版本：1
	VersionType *int32 `json:"version_type,omitempty"`

	// 审核角色
	ApproveRole *string `json:"approve_role,omitempty"`

	// request offset, from 0
	Offset *int32 `json:"offset,omitempty"`

	// request limit size
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListPlaybookVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPlaybookVersionsRequest struct{}"
	}

	return strings.Join([]string{"ListPlaybookVersionsRequest", string(data)}, " ")
}
