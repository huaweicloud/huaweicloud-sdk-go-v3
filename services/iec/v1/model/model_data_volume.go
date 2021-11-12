package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据盘参数
type DataVolume struct {
	// 数据盘大小，容量单位为GB，输入大小范围为[1,500]。

	Size int32 `json:"size"`
	// 边缘实例数据盘对应的磁盘类型，需要与站点所提供的磁盘类型相匹配。

	VolumeType string `json:"volume_type"`
}

func (o DataVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataVolume struct{}"
	}

	return strings.Join([]string{"DataVolume", string(data)}, " ")
}
