package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CompareConfigurationRequestBody struct {

	// 需要进行比较的源参数模板ID。
	SourceConfigurationId string `json:"source_configuration_id"`

	// 需要进行比较的目标参数模板ID。
	TargetConfigurationId string `json:"target_configuration_id"`
}

func (o CompareConfigurationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareConfigurationRequestBody struct{}"
	}

	return strings.Join([]string{"CompareConfigurationRequestBody", string(data)}, " ")
}
