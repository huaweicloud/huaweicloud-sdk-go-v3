package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHttpsPasswordSettingRequest Request Object
type UpdateHttpsPasswordSettingRequest struct {
	Body *UpdateHttpsPasswordSetting `json:"body,omitempty"`
}

func (o UpdateHttpsPasswordSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHttpsPasswordSettingRequest struct{}"
	}

	return strings.Join([]string{"UpdateHttpsPasswordSettingRequest", string(data)}, " ")
}
