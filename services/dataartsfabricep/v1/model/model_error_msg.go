package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ErrorMsg 错误描述
type ErrorMsg struct {
}

func (o ErrorMsg) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrorMsg struct{}"
	}

	return strings.Join([]string{"ErrorMsg", string(data)}, " ")
}
