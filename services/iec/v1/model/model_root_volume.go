package model

import (
	"encoding/json"

	"strings"
)

// 系统盘参数
type RootVolume struct {
	// 系统盘大小，容量单位为GB，输入大小范围为[40,100]。

	Size int32 `json:"size"`
	// 边缘实例系统盘对应的磁盘类型，需要与站点所提供的磁盘类型相匹配。

	VolumeType string `json:"volume_type"`
}

func (o RootVolume) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RootVolume struct{}"
	}

	return strings.Join([]string{"RootVolume", string(data)}, " ")
}
