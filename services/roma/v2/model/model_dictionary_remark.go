package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 字典描述
type DictionaryRemark struct {
}

func (o DictionaryRemark) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DictionaryRemark struct{}"
	}

	return strings.Join([]string{"DictionaryRemark", string(data)}, " ")
}
