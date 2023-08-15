package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Level 告警级别, 1为紧急，2为重要，3为次要，4为提示
type Level struct {
}

func (o Level) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Level struct{}"
	}

	return strings.Join([]string{"Level", string(data)}, " ")
}
