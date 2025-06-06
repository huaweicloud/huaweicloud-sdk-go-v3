package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyId 策略ID
type PolicyId struct {
}

func (o PolicyId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyId struct{}"
	}

	return strings.Join([]string{"PolicyId", string(data)}, " ")
}
