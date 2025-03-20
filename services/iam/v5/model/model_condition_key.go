package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConditionKey 条件键的名称。
type ConditionKey struct {
}

func (o ConditionKey) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConditionKey struct{}"
	}

	return strings.Join([]string{"ConditionKey", string(data)}, " ")
}
