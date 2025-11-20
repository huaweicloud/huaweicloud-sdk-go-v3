package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RelatedDnVo struct {

	// DN实例id。
	InstanceId *string `json:"instance_id,omitempty"`

	// DN实例名称。
	InstanceName *string `json:"instance_name,omitempty"`

	// DN实例状态。
	InstanceStatus *string `json:"instance_status,omitempty"`

	// 磁盘大小。
	DataVolumeSize float32 `json:"data_volume_size,omitempty"`

	// 版本。
	Version *string `json:"version,omitempty"`

	// 引擎名称。
	EngineName *string `json:"engine_name,omitempty"`
}

func (o RelatedDnVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RelatedDnVo struct{}"
	}

	return strings.Join([]string{"RelatedDnVo", string(data)}, " ")
}
