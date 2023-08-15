package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DictionaryOrder 字典排序，值越小顺序越靠前
type DictionaryOrder struct {
}

func (o DictionaryOrder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DictionaryOrder struct{}"
	}

	return strings.Join([]string{"DictionaryOrder", string(data)}, " ")
}
