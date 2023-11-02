package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ActionInfo 剧本流程动作信息
type ActionInfo struct {

	// 剧本流程动作ID
	Id *string `json:"id,omitempty"`

	// 流程动作名称
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 流程动作类型
	ActionType *string `json:"action_type,omitempty"`

	// 流程ID
	ActionId *string `json:"action_id,omitempty"`

	// 剧本ID
	PlaybookId *string `json:"playbook_id,omitempty"`

	// 剧本版本ID
	PlaybookVersionId *string `json:"playbook_version_id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`
}

func (o ActionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionInfo struct{}"
	}

	return strings.Join([]string{"ActionInfo", string(data)}, " ")
}
