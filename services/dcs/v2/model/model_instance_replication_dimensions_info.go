package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 监控指标维度对象信息
type InstanceReplicationDimensionsInfo struct {
	// 监控维度名称

	Name *string `json:"name,omitempty"`
	// 维度取值

	Value *string `json:"value,omitempty"`
}

func (o InstanceReplicationDimensionsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceReplicationDimensionsInfo struct{}"
	}

	return strings.Join([]string{"InstanceReplicationDimensionsInfo", string(data)}, " ")
}
