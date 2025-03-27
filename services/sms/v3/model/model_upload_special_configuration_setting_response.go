package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadSpecialConfigurationSettingResponse Response Object
type UploadSpecialConfigurationSettingResponse struct {

	// 上传相关配置成功
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UploadSpecialConfigurationSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadSpecialConfigurationSettingResponse struct{}"
	}

	return strings.Join([]string{"UploadSpecialConfigurationSettingResponse", string(data)}, " ")
}
