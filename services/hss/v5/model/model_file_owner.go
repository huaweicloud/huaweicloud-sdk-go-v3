package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FileOwner **参数解释**： 文件属性 **取值范围**： 字符长度0-64位
type FileOwner struct {
}

func (o FileOwner) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileOwner struct{}"
	}

	return strings.Join([]string{"FileOwner", string(data)}, " ")
}
