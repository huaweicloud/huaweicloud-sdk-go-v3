package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateInstanceConfigurationRequestBody struct {
	Values *UpdateInstanceConfigurationValuesOption `json:"values"`
}

func (o UpdateInstanceConfigurationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceConfigurationRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateInstanceConfigurationRequestBody", string(data)}, " ")
}
