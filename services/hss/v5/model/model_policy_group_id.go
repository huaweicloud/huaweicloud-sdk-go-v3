package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyGroupId **参数解释**： 策略组ID **取值范围**： 字符长度36-64位
type PolicyGroupId struct {
}

func (o PolicyGroupId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyGroupId struct{}"
	}

	return strings.Join([]string{"PolicyGroupId", string(data)}, " ")
}
