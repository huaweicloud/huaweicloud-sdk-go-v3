package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OsTypeRes **参数解释**： 操作系统类型 **取值范围**： - Linux：Linux - Windows：Windows
type OsTypeRes struct {
}

func (o OsTypeRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OsTypeRes struct{}"
	}

	return strings.Join([]string{"OsTypeRes", string(data)}, " ")
}
