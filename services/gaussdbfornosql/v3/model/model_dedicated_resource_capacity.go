package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 专属资源信息容量信息。
type DedicatedResourceCapacity struct {
	// CPU核数。

	Vcpus int32 `json:"vcpus"`
	// 内存大小，单位GB。

	Ram int32 `json:"ram"`
	// 存储大小，单位GB

	Volume int32 `json:"volume"`
}

func (o DedicatedResourceCapacity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DedicatedResourceCapacity struct{}"
	}

	return strings.Join([]string{"DedicatedResourceCapacity", string(data)}, " ")
}
