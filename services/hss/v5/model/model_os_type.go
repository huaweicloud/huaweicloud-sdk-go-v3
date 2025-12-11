package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OsType **参数解释**： 操作系统类型 **取值范围**： - Linux：Linux - Windows：Windows
type OsType struct {
}

func (o OsType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OsType struct{}"
	}

	return strings.Join([]string{"OsType", string(data)}, " ")
}
