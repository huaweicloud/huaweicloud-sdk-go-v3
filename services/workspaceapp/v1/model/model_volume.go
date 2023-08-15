package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Volume 磁盘
type Volume struct {
	Type *VolumeType `json:"type"`

	// 磁盘容量，单位GB，数值约束为10的倍数 * `系统盘` minLength: 10，maxLength: 1024 * `数据盘` minLength: 10，maxLength: 32768
	Size int32 `json:"size"`

	// 云服务器系统盘对应的存储池的ID
	ClusterId *string `json:"cluster_id,omitempty"`
}

func (o Volume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Volume struct{}"
	}

	return strings.Join([]string{"Volume", string(data)}, " ")
}
