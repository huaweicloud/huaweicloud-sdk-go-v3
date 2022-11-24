package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// notebook规格信息
type FlavorInfo struct {

	// notebook占用的cpu,0.1核为100m,单位为\"C\"
	Cpu float32 `json:"cpu"`

	// notebook占用的gpu，0.1为使用单卡10%，1为占满单个显卡，1+为使用多个显卡
	Gpu float32 `json:"gpu"`

	// notebook占用的内存,单位为\"G\"
	Memory float32 `json:"memory"`
}

func (o FlavorInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorInfo struct{}"
	}

	return strings.Join([]string{"FlavorInfo", string(data)}, " ")
}
