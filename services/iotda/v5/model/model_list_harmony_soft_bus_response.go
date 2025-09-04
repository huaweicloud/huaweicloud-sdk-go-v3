package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHarmonySoftBusResponse Response Object
type ListHarmonySoftBusResponse struct {

	// 设备组信息列表。
	HarmonySoftBuses *[]HarmonySoftBusResponseDto `json:"harmony_soft_buses,omitempty"`

	Page           *Page `json:"page,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListHarmonySoftBusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHarmonySoftBusResponse struct{}"
	}

	return strings.Join([]string{"ListHarmonySoftBusResponse", string(data)}, " ")
}
