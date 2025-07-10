package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VolumeAddInfo 磁盘信息。
type VolumeAddInfo struct {

	// 磁盘记录ID，删除或者扩容磁盘时必选。
	Id *string `json:"id,omitempty"`

	// 桌面数据盘对应的磁盘类型，需要与系统所提供的磁盘类型相匹配。 - SAS：高IO。 - SSD：超高IO。
	Type string `json:"type"`

	// 磁盘容量，单位GB。
	Size int32 `json:"size"`
}

func (o VolumeAddInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeAddInfo struct{}"
	}

	return strings.Join([]string{"VolumeAddInfo", string(data)}, " ")
}
