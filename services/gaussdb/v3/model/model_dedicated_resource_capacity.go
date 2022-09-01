package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DedicatedResourceCapacity struct {

	// 内存大小，单位GB
	Ram *int32 `json:"ram,omitempty" xml:"ram"`

	// 磁盘容量，单位GB
	Volume *int64 `json:"volume,omitempty" xml:"volume"`

	// cpu核数
	Vcpus *int32 `json:"vcpus,omitempty" xml:"vcpus"`
}

func (o DedicatedResourceCapacity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DedicatedResourceCapacity struct{}"
	}

	return strings.Join([]string{"DedicatedResourceCapacity", string(data)}, " ")
}
