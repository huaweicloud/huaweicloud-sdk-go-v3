package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MaskName 屏蔽规则名称，只能为字母、数字、汉字、-、_，最大长度为64
type MaskName struct {
}

func (o MaskName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MaskName struct{}"
	}

	return strings.Join([]string{"MaskName", string(data)}, " ")
}
