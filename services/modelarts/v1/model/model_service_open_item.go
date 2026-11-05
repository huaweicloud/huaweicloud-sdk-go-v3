package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServiceOpenItem struct {

	// **参数解释**：服务编码。 **约束限制**：不涉及。 **取值范围**：长度限制64个字符。 **默认取值**：不涉及。
	ServiceName string `json:"service_name"`
}

func (o ServiceOpenItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceOpenItem struct{}"
	}

	return strings.Join([]string{"ServiceOpenItem", string(data)}, " ")
}
