package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MessagePropertyList struct {

	// 属性名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 属性值。
	Value *string `json:"value,omitempty" xml:"value"`
}

func (o MessagePropertyList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MessagePropertyList struct{}"
	}

	return strings.Join([]string{"MessagePropertyList", string(data)}, " ")
}
