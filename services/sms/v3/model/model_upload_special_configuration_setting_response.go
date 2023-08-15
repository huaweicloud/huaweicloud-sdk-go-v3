package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadSpecialConfigurationSettingResponse Response Object
type UploadSpecialConfigurationSettingResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UploadSpecialConfigurationSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadSpecialConfigurationSettingResponse struct{}"
	}

	return strings.Join([]string{"UploadSpecialConfigurationSettingResponse", string(data)}, " ")
}
