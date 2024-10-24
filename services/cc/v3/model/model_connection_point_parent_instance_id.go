package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConnectionPointParentInstanceId 连接点的实例的父资源ID。
type ConnectionPointParentInstanceId struct {

	// 连接点的实例的父资源ID。
	ParentInstanceId *string `json:"parent_instance_id,omitempty"`
}

func (o ConnectionPointParentInstanceId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionPointParentInstanceId struct{}"
	}

	return strings.Join([]string{"ConnectionPointParentInstanceId", string(data)}, " ")
}
