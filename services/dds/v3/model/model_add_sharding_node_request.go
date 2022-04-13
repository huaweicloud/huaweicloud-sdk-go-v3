package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AddShardingNodeRequest struct {
	// 实例ID，可以调用“查询实例列表和详情”接口获取。如果未申请实例，可以调用“创建实例”接口创建。

	InstanceId string `json:"instance_id"`

	Body *EnlargeInstanceRequestBody `json:"body,omitempty"`
}

func (o AddShardingNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddShardingNodeRequest struct{}"
	}

	return strings.Join([]string{"AddShardingNodeRequest", string(data)}, " ")
}
