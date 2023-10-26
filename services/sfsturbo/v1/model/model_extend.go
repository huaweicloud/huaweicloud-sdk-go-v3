package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Extend extend对象
type Extend struct {

	// 扩容后文件系统的新容量，以GiB为单位。  通用型-标准型和性能型，取值范围500~32768 GiB，扩容步长大于等于100 GiB。  通用型-标准型增强版和性能型增强版。容量范围是10240~327680 GiB。扩容步长大于等于100 GiB。  HPC型文件系统，容量范围是3686~1048576 GiB。HPC型文件系统的容量必须为1.2TiB的倍数，扩容步长需要大于等于1.2TiB，需要将目标容量换算为GiB后需要向下取整。如4.8TiB->4915GiB，8.4TiB->8601GiB。  HPC缓存型文件系统，容量范围是4096~1048576 GiB。扩容步长均为1TiB。需要将目标容量换算为GiB。
	NewSize int32 `json:"new_size"`

	// 带宽的目标值，单位：GB。仅HPC缓存型支持带宽扩缩。 支持的带宽值为：\"2G\"、\"4G\"、\"8G\"、\"16G\"、\"24G\"、\"32G\"、\"48G\"。
	NewBandwidth *int64 `json:"new_bandwidth,omitempty"`

	BssParam *BssInfoExtend `json:"bss_param,omitempty"`
}

func (o Extend) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Extend struct{}"
	}

	return strings.Join([]string{"Extend", string(data)}, " ")
}
