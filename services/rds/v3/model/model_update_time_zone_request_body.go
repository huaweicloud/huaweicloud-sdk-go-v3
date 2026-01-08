package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateTimeZoneRequestBody struct {

	// 时区
	TimeZone string `json:"time_zone"`
}

func (o UpdateTimeZoneRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTimeZoneRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateTimeZoneRequestBody", string(data)}, " ")
}
