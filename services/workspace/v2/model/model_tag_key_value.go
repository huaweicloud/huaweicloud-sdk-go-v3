package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagKeyValue 镜像标签键值。
type TagKeyValue struct {

	// 标签的键。
	Key *string `json:"key,omitempty"`

	// 标签的值。
	Value *string `json:"value,omitempty"`
}

func (o TagKeyValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagKeyValue struct{}"
	}

	return strings.Join([]string{"TagKeyValue", string(data)}, " ")
}
