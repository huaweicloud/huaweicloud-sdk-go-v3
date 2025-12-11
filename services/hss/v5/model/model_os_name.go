package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OsName **参数解释**: 操作系统名称 **取值范围**: 字符长度0-128位
type OsName struct {
}

func (o OsName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OsName struct{}"
	}

	return strings.Join([]string{"OsName", string(data)}, " ")
}
