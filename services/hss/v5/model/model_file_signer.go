package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FileSigner **参数解释**： 文件签名 **取值范围**： 字符长度1-128位
type FileSigner struct {
}

func (o FileSigner) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileSigner struct{}"
	}

	return strings.Join([]string{"FileSigner", string(data)}, " ")
}
