package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEdgeNodeDetailRequest Request Object
type ShowEdgeNodeDetailRequest struct {

	// 边缘节点ID
	NodeId string `json:"node_id"`

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`
}

func (o ShowEdgeNodeDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEdgeNodeDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowEdgeNodeDetailRequest", string(data)}, " ")
}
