package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DictionaryRemark 字典描述
type DictionaryRemark struct {
}

func (o DictionaryRemark) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DictionaryRemark struct{}"
	}

	return strings.Join([]string{"DictionaryRemark", string(data)}, " ")
}
