package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HwcEcsFlavor 弹性云服务器规格信息。
type HwcEcsFlavor struct {

	// 云服务器规格ID。
	Id string `json:"id"`

	// 云服务器规格名称。
	Name string `json:"name"`

	// 该云服务器规格对应要求系统盘大小，0为不限制。
	Disk *string `json:"disk,omitempty"`

	// 该云服务器规格对应的CPU核数。
	Vcpus *string `json:"vcpus,omitempty"`

	// 该云服务器规格对应的内存大小，单位为MB。
	Ram *string `json:"ram,omitempty"`
}

func (o HwcEcsFlavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HwcEcsFlavor struct{}"
	}

	return strings.Join([]string{"HwcEcsFlavor", string(data)}, " ")
}
