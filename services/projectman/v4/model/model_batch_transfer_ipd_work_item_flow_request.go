package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchTransferIpdWorkItemFlowRequest Request Object
type BatchTransferIpdWorkItemFlowRequest struct {

	// 项目32位ID，项目唯一标识。通过查询IPD项目列表获取，响应消息体中的id字段的值就是项目ID。
	ProjectId string `json:"project_id"`

	// **参数解释**： 是否覆盖对应字段。 **约束限制**： 不涉及 **取值范围**： true:本开关开启时，当前弹窗的相应字段值将覆盖全部所选工作项的对应字段值。 false:本开关关闭时，除「当前责任人」之外，所选工作项的对应字段如果已经有值，将保持原状，不会被当前弹窗的相应字段值覆盖。 **默认取值**： false。
	IsRecover *bool `json:"is_recover,omitempty"`

	Body *WorkItemFlowVo `json:"body,omitempty"`
}

func (o BatchTransferIpdWorkItemFlowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchTransferIpdWorkItemFlowRequest struct{}"
	}

	return strings.Join([]string{"BatchTransferIpdWorkItemFlowRequest", string(data)}, " ")
}
