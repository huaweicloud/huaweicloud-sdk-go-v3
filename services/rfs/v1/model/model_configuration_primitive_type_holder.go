package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigurationPrimitiveTypeHolder struct {
	Configuration *ConfigurationPrimitiveTypeHolderConfiguration `json:"configuration,omitempty"`
}

func (o ConfigurationPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"ConfigurationPrimitiveTypeHolder", string(data)}, " ")
}
