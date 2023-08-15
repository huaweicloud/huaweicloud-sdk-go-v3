package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VolumeResp 如果规格为固定存储容量规格，则该属性为规格典配的存储容量信息，如果为弹性存储规格，则该属性为null。
type VolumeResp struct {

	// 磁盘类型，仅支持SSD。
	Type string `json:"type"`

	// 磁盘可用容量。
	Size int32 `json:"size"`
}

func (o VolumeResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeResp struct{}"
	}

	return strings.Join([]string{"VolumeResp", string(data)}, " ")
}
