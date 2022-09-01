package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateEdgeNodeDeviceRequest struct {

	// 节点ID
	NodeId string `json:"node_id" xml:"node_id"`

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty" xml:"ief-instance-id"`

	Body *Devices `json:"body,omitempty" xml:"body"`
}

func (o UpdateEdgeNodeDeviceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEdgeNodeDeviceRequest struct{}"
	}

	return strings.Join([]string{"UpdateEdgeNodeDeviceRequest", string(data)}, " ")
}
