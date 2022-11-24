package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type FlavorInstance struct {

	// 边缘实例规格的ID。
	Id *string `json:"id,omitempty"`

	// 边缘实例规格的名称。
	Name *string `json:"name,omitempty"`

	// 边缘实例规格对应要求系统盘大小。  当前未使用该参数，缺省值为0。
	Disk *string `json:"disk,omitempty"`

	// 边缘实例规格对应的内存大小，单位为MB。
	Ram *string `json:"ram,omitempty"`

	// 边缘实例规格对应的CPU核数。
	Vcpus *string `json:"vcpus,omitempty"`
}

func (o FlavorInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorInstance struct{}"
	}

	return strings.Join([]string{"FlavorInstance", string(data)}, " ")
}
