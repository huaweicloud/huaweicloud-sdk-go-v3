package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 专属资源信息列表。
type ListDedicatedResourceResult struct {
	// 专属资源ID。

	Id string `json:"id"`
	// 专属资源的名称。

	ResourceName string `json:"resource_name"`
	// 引擎名称。

	EngineName string `json:"engine_name"`
	// 可用区信息。

	AvailabilityZone string `json:"availability_zone"`
	// 专属资源的计算架构。

	Architecture string `json:"architecture"`

	Capacity *DedicatedResourceCapacity `json:"capacity"`
	// 专属资源的状态信息。

	Status string `json:"status"`
}

func (o ListDedicatedResourceResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDedicatedResourceResult struct{}"
	}

	return strings.Join([]string{"ListDedicatedResourceResult", string(data)}, " ")
}
