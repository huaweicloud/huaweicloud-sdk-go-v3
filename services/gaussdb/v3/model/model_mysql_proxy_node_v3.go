package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MysqlProxyNodeV3 struct {

	// 节点id。
	Id *string `json:"id,omitempty" xml:"id"`

	// 实例id。
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id"`

	// 节点状态。
	Status *string `json:"status,omitempty" xml:"status"`

	// 节点名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 节点读写分离读权重。
	Weight *int32 `json:"weight,omitempty" xml:"weight"`

	// 可用区信息。
	AvailabilityZone *[]MysqlProxyAvailable `json:"availability_zone,omitempty" xml:"availability_zone"`
}

func (o MysqlProxyNodeV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlProxyNodeV3 struct{}"
	}

	return strings.Join([]string{"MysqlProxyNodeV3", string(data)}, " ")
}
