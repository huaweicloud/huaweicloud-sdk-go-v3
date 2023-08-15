package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AdmetPropertyDictWithCustom 分子ADMET属性字典(包含自定义属性)
type AdmetPropertyDictWithCustom struct {
}

func (o AdmetPropertyDictWithCustom) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdmetPropertyDictWithCustom struct{}"
	}

	return strings.Join([]string{"AdmetPropertyDictWithCustom", string(data)}, " ")
}
