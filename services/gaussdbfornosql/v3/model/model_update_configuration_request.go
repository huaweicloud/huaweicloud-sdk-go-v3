package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateConfigurationRequest struct {

	// 参数模板ID。
	ConfigId string `json:"config_id" xml:"config_id"`

	Body *UpdateConfigurationRequestBody `json:"body,omitempty" xml:"body"`
}

func (o UpdateConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigurationRequest struct{}"
	}

	return strings.Join([]string{"UpdateConfigurationRequest", string(data)}, " ")
}
