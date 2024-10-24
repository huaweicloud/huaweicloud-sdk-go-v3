package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConnectionPointInstanceId 连接点的实例ID。
type ConnectionPointInstanceId struct {

	// 连接点的实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ConnectionPointInstanceId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionPointInstanceId struct{}"
	}

	return strings.Join([]string{"ConnectionPointInstanceId", string(data)}, " ")
}
