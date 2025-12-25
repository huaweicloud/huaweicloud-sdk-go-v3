package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifierName 修改者名称
type ModifierName struct {
}

func (o ModifierName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifierName struct{}"
	}

	return strings.Join([]string{"ModifierName", string(data)}, " ")
}
