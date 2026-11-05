package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServiceStatusItem struct {

	// **参数解释**：服务编码。 **取值范围**：长度限制64个字符。
	ServiceName string `json:"service_name"`

	// **参数解释**：服务开通状态。 **取值范围**：不涉及。 - true：已开通。 - false：未开通。
	Status bool `json:"status"`
}

func (o ServiceStatusItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceStatusItem struct{}"
	}

	return strings.Join([]string{"ServiceStatusItem", string(data)}, " ")
}
