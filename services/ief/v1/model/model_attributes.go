package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Attributes 节点属性
type Attributes struct {

	// 节点属性的key值，长度取值范围为1~128， 仅允许大小写英文字母、数字、下划线、中划线
	Key *string `json:"key,omitempty"`

	// 节点属性的value值，长度取值范围为1~256， 仅允许大小写英文字母、数字、下划线、中划线
	Value *string `json:"value,omitempty"`
}

func (o Attributes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Attributes struct{}"
	}

	return strings.Join([]string{"Attributes", string(data)}, " ")
}
