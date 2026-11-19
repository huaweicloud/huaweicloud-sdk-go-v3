package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppNameRes **参数解释**: 软件名称 **取值范围**: 字符长度0-128位
type AppNameRes struct {
}

func (o AppNameRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppNameRes struct{}"
	}

	return strings.Join([]string{"AppNameRes", string(data)}, " ")
}
