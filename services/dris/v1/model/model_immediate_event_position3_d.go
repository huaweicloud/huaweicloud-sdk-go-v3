package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImmediateEventPosition3D 位置信息。
type ImmediateEventPosition3D struct {

	// **参数说明**：定义纬度数值，北纬为正，南纬为负，单位°，精度小数点后7位。
	Lat float32 `json:"lat"`

	// **参数说明**：定义经度数值。东经为正，西经为负，单位°，精度小数点后7位。
	Lon float32 `json:"lon"`

	// **参数说明**：定义海拔高程，可选，单位为分米。
	Ele float32 `json:"ele,omitempty"`
}

func (o ImmediateEventPosition3D) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImmediateEventPosition3D struct{}"
	}

	return strings.Join([]string{"ImmediateEventPosition3D", string(data)}, " ")
}
