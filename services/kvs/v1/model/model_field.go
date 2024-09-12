package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Field struct {

	// 字段名。
	Name string `bson:"name"`

	// bool值预留无意义。
	Order *bool `bson:"order,omitempty"`
}

func (o Field) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Field struct{}"
	}

	return strings.Join([]string{"Field", string(data)}, " ")
}
