package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyGroupName 策略组名
type PolicyGroupName struct {
}

func (o PolicyGroupName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyGroupName struct{}"
	}

	return strings.Join([]string{"PolicyGroupName", string(data)}, " ")
}
