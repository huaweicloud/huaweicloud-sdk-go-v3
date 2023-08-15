package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetConfigurationRequest Request Object
type ResetConfigurationRequest struct {

	// 需重置的参数模板ID。
	ConfigId string `json:"config_id"`
}

func (o ResetConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetConfigurationRequest struct{}"
	}

	return strings.Join([]string{"ResetConfigurationRequest", string(data)}, " ")
}
