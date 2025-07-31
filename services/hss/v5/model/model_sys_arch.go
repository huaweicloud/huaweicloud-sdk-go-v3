package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SysArch **参数解释**： 系统CPU架构 **取值范围**： 字符长度1-64位
type SysArch struct {
}

func (o SysArch) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SysArch struct{}"
	}

	return strings.Join([]string{"SysArch", string(data)}, " ")
}
