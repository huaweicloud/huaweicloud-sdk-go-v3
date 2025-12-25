package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPlaybookTopologyRequest Request Object
type ShowPlaybookTopologyRequest struct {

	// **参数解释：** 内容类型 - application/json    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json  **默认取值：** application/json
	ContentType string `json:"content-type"`

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 剧本实例ID
	InstanceId string `json:"instance_id"`
}

func (o ShowPlaybookTopologyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPlaybookTopologyRequest struct{}"
	}

	return strings.Join([]string{"ShowPlaybookTopologyRequest", string(data)}, " ")
}
