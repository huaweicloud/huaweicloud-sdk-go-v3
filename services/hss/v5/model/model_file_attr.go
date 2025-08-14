package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FileAttr **参数解释**： 文件属性 **取值范围**： 字符长度1-256位
type FileAttr struct {
}

func (o FileAttr) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileAttr struct{}"
	}

	return strings.Join([]string{"FileAttr", string(data)}, " ")
}
