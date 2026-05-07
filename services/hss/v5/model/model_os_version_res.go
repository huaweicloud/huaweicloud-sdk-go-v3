package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OsVersionRes **参数解释**: 操作系统版本 **取值范围**: 字符长度1-64位
type OsVersionRes struct {
}

func (o OsVersionRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OsVersionRes struct{}"
	}

	return strings.Join([]string{"OsVersionRes", string(data)}, " ")
}
