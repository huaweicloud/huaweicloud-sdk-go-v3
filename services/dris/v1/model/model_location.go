package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// **参数说明**：位置信息。
type Location struct {

	// **参数说明**：定义纬度数值，北纬为正，南纬为负，单位°，精度小数点后7位。
	Lat float32 `json:"lat,omitempty"`

	// **参数说明**：定义经度数值。东经为正，西经为负。单位°，精度小数点后7位。
	Lon float32 `json:"lon,omitempty"`
}

func (o Location) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Location struct{}"
	}

	return strings.Join([]string{"Location", string(data)}, " ")
}
