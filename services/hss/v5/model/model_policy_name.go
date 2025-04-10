package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyName 策略名称
type PolicyName struct {
}

func (o PolicyName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyName struct{}"
	}

	return strings.Join([]string{"PolicyName", string(data)}, " ")
}
