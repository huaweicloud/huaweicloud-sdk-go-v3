package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConfigurationDetailRequest Request Object
type ShowConfigurationDetailRequest struct {

	// 参数模板ID。
	ConfigId string `json:"config_id"`
}

func (o ShowConfigurationDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigurationDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowConfigurationDetailRequest", string(data)}, " ")
}
