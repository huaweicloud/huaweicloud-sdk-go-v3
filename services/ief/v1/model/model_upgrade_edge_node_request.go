package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeEdgeNodeRequest Request Object
type UpgradeEdgeNodeRequest struct {

	// 边缘节点ID
	NodeId string `json:"node_id"`

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`
}

func (o UpgradeEdgeNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeEdgeNodeRequest struct{}"
	}

	return strings.Join([]string{"UpgradeEdgeNodeRequest", string(data)}, " ")
}
