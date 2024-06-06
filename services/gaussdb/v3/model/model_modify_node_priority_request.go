package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyNodePriorityRequest Request Object
type ModifyNodePriorityRequest struct {

	// 内容类型。  取值：application/json。
	ContentType string `json:"Content-Type"`

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 节点ID。
	NodeId string `json:"node_id"`

	Body *ModifyNodePriorityRequestBody `json:"body,omitempty"`
}

func (o ModifyNodePriorityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyNodePriorityRequest struct{}"
	}

	return strings.Join([]string{"ModifyNodePriorityRequest", string(data)}, " ")
}
