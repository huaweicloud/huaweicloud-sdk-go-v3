package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAopWorkflowInstanceRequest Request Object
type ShowAopWorkflowInstanceRequest struct {

	// **参数解释：** 内容类型 - application/json;charset=UTF-8    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json;charset=UTF-8  **默认取值：** 不涉及
	ContentType string `json:"Content-Type"`

	// **参数解释**: 工作空间ID **取值范围**: 不涉及
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**: 流程实例的ID **约束限制**: 不涉及
	InstanceId string `json:"instance_id"`

	// **参数解释**: 是否查询流程拓扑图的信息 - true 查询 - false 不查询  **约束限制**: 不涉及               **取值范围**: - true - false  **默认值**:  false
	ShowTopology *bool `json:"show_topology,omitempty"`
}

func (o ShowAopWorkflowInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAopWorkflowInstanceRequest struct{}"
	}

	return strings.Join([]string{"ShowAopWorkflowInstanceRequest", string(data)}, " ")
}
