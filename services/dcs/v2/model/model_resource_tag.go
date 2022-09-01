package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceTag struct {

	// 标签键 - 不能为空。 - 对于同一资源键值唯一。 - 长度不超过36个字符。 - 标签的键必须唯一且输入不能为空。
	Key string `json:"key" xml:"key"`

	// 标签值 - action为create时必选。action为delete时非必选。 - 长度不超过43个字符。
	Value *string `json:"value,omitempty" xml:"value"`
}

func (o ResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTag struct{}"
	}

	return strings.Join([]string{"ResourceTag", string(data)}, " ")
}
