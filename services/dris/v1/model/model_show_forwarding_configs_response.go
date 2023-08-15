package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowForwardingConfigsResponse Response Object
type ShowForwardingConfigsResponse struct {

	// **参数说明**：满足查询条件的记录总数。
	Count *int32 `json:"count,omitempty"`

	// **参数说明**：转发配置的列表。
	ForwardingConfigs *[]ForwardingConfig `json:"forwarding_configs,omitempty"`
	HttpStatusCode    int                 `json:"-"`
}

func (o ShowForwardingConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowForwardingConfigsResponse struct{}"
	}

	return strings.Join([]string{"ShowForwardingConfigsResponse", string(data)}, " ")
}
