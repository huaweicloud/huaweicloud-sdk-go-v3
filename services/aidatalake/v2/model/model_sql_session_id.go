package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SqlSessionId **参数解释**：Session的ID。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。
type SqlSessionId struct {
}

func (o SqlSessionId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlSessionId struct{}"
	}

	return strings.Join([]string{"SqlSessionId", string(data)}, " ")
}
