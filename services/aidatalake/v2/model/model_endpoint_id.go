package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EndpointId **参数解释**：端点ID。 **取值范围**：长度为1~64个字符，支持大小写英文字母、数字、连字符。
type EndpointId struct {
}

func (o EndpointId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointId struct{}"
	}

	return strings.Join([]string{"EndpointId", string(data)}, " ")
}
