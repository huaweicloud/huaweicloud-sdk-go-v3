package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DiffConfigurationRequest struct {

	// 需要进行比较的参数模板ID。
	SourceConfigurationId string `json:"source_configuration_id"`

	// 需要进行比较的参数模板ID。
	TargetConfigurationId string `json:"target_configuration_id"`
}

func (o DiffConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiffConfigurationRequest struct{}"
	}

	return strings.Join([]string{"DiffConfigurationRequest", string(data)}, " ")
}
