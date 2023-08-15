package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandReplicasetNodeRequest Request Object
type ExpandReplicasetNodeRequest struct {

	// 实例ID，可以调用“查询实例列表和详情”接口获取。如果未申请实例，可以调用“创建实例”接口创建。
	InstanceId string `json:"instance_id"`

	Body *EnlargeReplicasetNodeRequestBody `json:"body,omitempty"`
}

func (o ExpandReplicasetNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandReplicasetNodeRequest struct{}"
	}

	return strings.Join([]string{"ExpandReplicasetNodeRequest", string(data)}, " ")
}
