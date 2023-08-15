package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConnectionStatisticsRequest Request Object
type ShowConnectionStatisticsRequest struct {

	// 实例ID，可以调用“查询实例列表和详情”接口获取。如果未申请实例，可以调用“创建实例”接口创建。
	InstanceId string `json:"instance_id"`

	// 节点ID。 - 如取空值，则默认查询实例下所有允许连接的节点的连接数信息。
	NodeId *string `json:"node_id,omitempty"`
}

func (o ShowConnectionStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConnectionStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ShowConnectionStatisticsRequest", string(data)}, " ")
}
