package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpsPasswordSettingRequest Request Object
type ShowHttpsPasswordSettingRequest struct {
}

func (o ShowHttpsPasswordSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpsPasswordSettingRequest struct{}"
	}

	return strings.Join([]string{"ShowHttpsPasswordSettingRequest", string(data)}, " ")
}
