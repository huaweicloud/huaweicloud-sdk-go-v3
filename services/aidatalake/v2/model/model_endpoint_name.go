package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EndpointName **参数解释**：Endpoint名称。 **约束限制**：不涉及。 **取值范围**：长度为1~128个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
type EndpointName struct {
}

func (o EndpointName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointName struct{}"
	}

	return strings.Join([]string{"EndpointName", string(data)}, " ")
}
