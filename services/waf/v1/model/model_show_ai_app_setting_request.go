package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAiAppSettingRequest Request Object
type ShowAiAppSettingRequest struct {
}

func (o ShowAiAppSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAiAppSettingRequest struct{}"
	}

	return strings.Join([]string{"ShowAiAppSettingRequest", string(data)}, " ")
}
