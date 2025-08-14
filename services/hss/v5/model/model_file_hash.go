package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FileHash **参数解释**： 文件哈希 **取值范围**： 字符长度1-256位
type FileHash struct {
}

func (o FileHash) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileHash struct{}"
	}

	return strings.Join([]string{"FileHash", string(data)}, " ")
}
