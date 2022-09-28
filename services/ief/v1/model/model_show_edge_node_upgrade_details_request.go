package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowEdgeNodeUpgradeDetailsRequest struct {

	// 边缘节点ID
	NodeId string `json:"node_id"`

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`
}

func (o ShowEdgeNodeUpgradeDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEdgeNodeUpgradeDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowEdgeNodeUpgradeDetailsRequest", string(data)}, " ")
}
