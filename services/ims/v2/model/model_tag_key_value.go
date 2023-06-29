package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagKeyValue 镜像标签
type TagKeyValue struct {

	// 标签的键
	Key string `json:"key"`

	// 标签的值
	Value string `json:"value"`
}

func (o TagKeyValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagKeyValue struct{}"
	}

	return strings.Join([]string{"TagKeyValue", string(data)}, " ")
}
