package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SimpleResourceNoUsed 简单配额资源，没有已使用值。
type SimpleResourceNoUsed struct {

	// 资源类别。 general_instances：普通桌面 Memory：内存 cores：CPU volumes：磁盘数量 volume_gigabytes：磁盘容量 gpu_instances：GPU桌面 deh：云办公主机 users：用户 policy_groups: 策略组 Cores: CPU(配额工具使用)
	Type *string `json:"type,omitempty"`

	// 配额数
	Quota *int32 `json:"quota,omitempty"`
}

func (o SimpleResourceNoUsed) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimpleResourceNoUsed struct{}"
	}

	return strings.Join([]string{"SimpleResourceNoUsed", string(data)}, " ")
}
