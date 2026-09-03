package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeVolumeChargeModeResponse Response Object
type ChangeVolumeChargeModeResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeVolumeChargeModeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeVolumeChargeModeResponse struct{}"
	}

	return strings.Join([]string{"ChangeVolumeChargeModeResponse", string(data)}, " ")
}
