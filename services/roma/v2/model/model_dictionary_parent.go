package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DictionaryParent 父字典编码,为空时代表自身就是最顶级字典
type DictionaryParent struct {
}

func (o DictionaryParent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DictionaryParent struct{}"
	}

	return strings.Join([]string{"DictionaryParent", string(data)}, " ")
}
