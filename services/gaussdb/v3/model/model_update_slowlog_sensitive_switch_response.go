package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSlowlogSensitiveSwitchResponse Response Object
type UpdateSlowlogSensitiveSwitchResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateSlowlogSensitiveSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSlowlogSensitiveSwitchResponse struct{}"
	}

	return strings.Join([]string{"UpdateSlowlogSensitiveSwitchResponse", string(data)}, " ")
}
