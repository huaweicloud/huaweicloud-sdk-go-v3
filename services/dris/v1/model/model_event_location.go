package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EventLocation struct {

	// **参数说明**：定义纬度数值，北纬为正，南纬为负，单位°，精度小数点后7位。
	Lat float32 `json:"lat,omitempty"`

	// **参数说明**：定义经度数值。东经为正，西经为负，单位°，精度小数点后7位。
	Lon float32 `json:"lon,omitempty"`
}

func (o EventLocation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventLocation struct{}"
	}

	return strings.Join([]string{"EventLocation", string(data)}, " ")
}
