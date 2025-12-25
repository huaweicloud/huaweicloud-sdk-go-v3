package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowComponentActionRequest Request Object
type ShowComponentActionRequest struct {

	// **参数解释**: 工作空间ID **取值范围**: 不涉及
	WorkspaceId string `json:"workspace_id"`

	// 插件id
	ComponentId string `json:"component_id"`

	// 插件执行函数id
	ActionId string `json:"action_id"`

	// 是否启用
	Enabled bool `json:"enabled"`
}

func (o ShowComponentActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowComponentActionRequest struct{}"
	}

	return strings.Join([]string{"ShowComponentActionRequest", string(data)}, " ")
}
