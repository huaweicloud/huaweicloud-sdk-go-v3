package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Position3D 位置信息。
type Position3D struct {

	// **参数说明**：定义纬度数值，北纬为正，南纬为负，单位°，精度小数点后7位。
	Lat *interface{} `json:"lat"`

	// **参数说明**：定义经度数值。东经为正，西经为负，单位°，精度小数点后7位。
	Lon *interface{} `json:"lon"`

	// **参数说明**：定义海拔高程，可选，单位为分米。
	Ele *interface{} `json:"ele,omitempty"`
}

func (o Position3D) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Position3D struct{}"
	}

	return strings.Join([]string{"Position3D", string(data)}, " ")
}
