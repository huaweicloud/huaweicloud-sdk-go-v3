package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateApplicationInstanceResponseConfigurationReqBody UpdateApplicationInstanceResponseConfiguration的请求体
type UpdateApplicationInstanceResponseConfigurationReqBody struct {
	ResponseConfig *ResponseConfigDto `json:"response_config"`
}

func (o UpdateApplicationInstanceResponseConfigurationReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApplicationInstanceResponseConfigurationReqBody struct{}"
	}

	return strings.Join([]string{"UpdateApplicationInstanceResponseConfigurationReqBody", string(data)}, " ")
}
