package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Target struct {
	// 实例ID，可以调用“查询实例列表和详情”接口获取。如果未申请实例，可以调用“创建实例”接口创建。

	InstanceId string `json:"instance_id"`
}

func (o Target) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Target struct{}"
	}

	return strings.Join([]string{"Target", string(data)}, " ")
}
