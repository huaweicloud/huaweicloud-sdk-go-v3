package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHarmonySoftBusResponse Response Object
type DeleteHarmonySoftBusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteHarmonySoftBusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHarmonySoftBusResponse struct{}"
	}

	return strings.Join([]string{"DeleteHarmonySoftBusResponse", string(data)}, " ")
}
