package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigurationParameterValues **参数解释**： 集群参数配置列表信息。 **取值范围**： 不涉及。
type ConfigurationParameterValues struct {

	// **参数解释**： 集群参数配置列表。 **取值范围**： 不涉及。
	Configurations []ConfigurationParameterValue `json:"configurations"`
}

func (o ConfigurationParameterValues) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationParameterValues struct{}"
	}

	return strings.Join([]string{"ConfigurationParameterValues", string(data)}, " ")
}
