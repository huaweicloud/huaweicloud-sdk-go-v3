package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyDescription 身份策略描述。
type PolicyDescription struct {
}

func (o PolicyDescription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyDescription struct{}"
	}

	return strings.Join([]string{"PolicyDescription", string(data)}, " ")
}
