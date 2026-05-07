package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OsNameRes **参数解释**: 操作系统名称 **取值范围**: 字符长度1-64位
type OsNameRes struct {
}

func (o OsNameRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OsNameRes struct{}"
	}

	return strings.Join([]string{"OsNameRes", string(data)}, " ")
}
