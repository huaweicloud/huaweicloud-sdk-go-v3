package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ErrorCode 错误码
type ErrorCode struct {
}

func (o ErrorCode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrorCode struct{}"
	}

	return strings.Join([]string{"ErrorCode", string(data)}, " ")
}
