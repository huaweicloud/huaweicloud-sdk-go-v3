package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomFeildRecord struct {

	// 自定义工作项属性
	Key *string `json:"key,omitempty" xml:"key"`

	// 自定义工作项名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 自定义工作项值
	Value *string `json:"value,omitempty" xml:"value"`
}

func (o CustomFeildRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomFeildRecord struct{}"
	}

	return strings.Join([]string{"CustomFeildRecord", string(data)}, " ")
}
