package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VolumeInfo 磁盘信息。
type VolumeInfo struct {

	// 磁盘类型。 - SATA：普通IO磁盘类型。 - SAS：高IO磁盘类型。 - SSD：超高IO磁盘类型。 - GPSSD：通用型SSD磁盘类型
	Type string `json:"type"`

	// 磁盘大小。单位为GB。
	Size int32 `json:"size"`

	// 磁盘数量。
	Count int32 `json:"count"`
}

func (o VolumeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeInfo struct{}"
	}

	return strings.Join([]string{"VolumeInfo", string(data)}, " ")
}
