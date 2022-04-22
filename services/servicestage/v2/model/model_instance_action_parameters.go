package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 操作参数，scale和rollback时提供。
type InstanceActionParameters struct {

	// 实例数，在scale操作时提供。
	Replica *int32 `json:"replica,omitempty"`

	// ECS ID列表，指定虚机扩容时部署的ECS主机。
	Hosts *[]string `json:"hosts,omitempty"`

	// 版本号，在rollback操作时提供，通过查询快照接口获取。
	Version *string `json:"version,omitempty"`
}

func (o InstanceActionParameters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceActionParameters struct{}"
	}

	return strings.Join([]string{"InstanceActionParameters", string(data)}, " ")
}
