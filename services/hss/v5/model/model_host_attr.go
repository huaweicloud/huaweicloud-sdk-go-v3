package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HostAttr **参数解释**： 主机属性 **取值范围**： 字符长度1-64位
type HostAttr struct {
}

func (o HostAttr) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostAttr struct{}"
	}

	return strings.Join([]string{"HostAttr", string(data)}, " ")
}
