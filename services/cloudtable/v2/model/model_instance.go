package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群的实例对象。
type Instance struct {
	// 集群所在可用区的ID。

	AvailabilityZone string `json:"availability_zone"`
	// CloudTable集群计算单元节点数目，至少为2。

	CuNum int32 `json:"cu_num"`
	// CloudTable集群Lemon节点数目。

	LemonNum *int32 `json:"lemon_num,omitempty"`
	// 集群所在网络信息。

	Nics []Nics `json:"nics"`
	// CloudTable集群TSD节点数目，至少为2。

	TsdNum int32 `json:"tsd_num"`
}

func (o Instance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Instance struct{}"
	}

	return strings.Join([]string{"Instance", string(data)}, " ")
}
