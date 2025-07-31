package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OsBit **参数解释**： 操作系统位数 **取值范围**： 字符长度1-64位
type OsBit struct {
}

func (o OsBit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OsBit struct{}"
	}

	return strings.Join([]string{"OsBit", string(data)}, " ")
}
