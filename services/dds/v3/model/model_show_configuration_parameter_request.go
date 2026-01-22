package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConfigurationParameterRequest Request Object
type ShowConfigurationParameterRequest struct {

	// **参数解释：** 参数模板ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ConfigId string `json:"config_id"`
}

func (o ShowConfigurationParameterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigurationParameterRequest struct{}"
	}

	return strings.Join([]string{"ShowConfigurationParameterRequest", string(data)}, " ")
}
