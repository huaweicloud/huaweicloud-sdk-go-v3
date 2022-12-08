package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据盘
type DataVolumes struct {

	// 对应的系统盘类型 SAS SATA SSD
	VolumeType string `json:"volume_type"`

	// 系统盘大小，容量单位为GB，输入大 小范围为[1-1024]
	Size string `json:"size"`

	// 磁盘产品信息
	ExtendParam string `json:"extend_param"`
}

func (o DataVolumes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataVolumes struct{}"
	}

	return strings.Join([]string{"DataVolumes", string(data)}, " ")
}
