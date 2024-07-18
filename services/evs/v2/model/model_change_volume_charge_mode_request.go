package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeVolumeChargeModeRequest Request Object
type ChangeVolumeChargeModeRequest struct {
	Body *ChangeVolumeChargeModeRequestBody `json:"body,omitempty"`
}

func (o ChangeVolumeChargeModeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeVolumeChargeModeRequest struct{}"
	}

	return strings.Join([]string{"ChangeVolumeChargeModeRequest", string(data)}, " ")
}
