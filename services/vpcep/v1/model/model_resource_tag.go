package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceTag struct {

	// 键。最大长度128个unicode字符。 key需要满足标签字符集规范。
	Key string `json:"key"`

	// 值。action为create时必选，每个值最大长度255个unicode字符， 删除时如果value有值按照key/value删除， 如果value没值，则按照key删除。 value需要满足标签字符集规范。
	Value *string `json:"value,omitempty"`
}

func (o ResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTag struct{}"
	}

	return strings.Join([]string{"ResourceTag", string(data)}, " ")
}
