package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EndpointIdAdmin **参数解释**：端点ID。 **取值范围**：长度为1~64个字符，支持大小写英文字母、数字、连字符。
type EndpointIdAdmin struct {
}

func (o EndpointIdAdmin) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointIdAdmin struct{}"
	}

	return strings.Join([]string{"EndpointIdAdmin", string(data)}, " ")
}
