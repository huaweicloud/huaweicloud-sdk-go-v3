package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagMatch 搜索字段，key为要匹配的字段，如resource_name等。value为匹配的值。
type TagMatch struct {

	// 键。第一期限定为resource_name，后续扩展。
	Key string `json:"key"`

	// 资源标签的值。
	Value *string `json:"value,omitempty"`
}

func (o TagMatch) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagMatch struct{}"
	}

	return strings.Join([]string{"TagMatch", string(data)}, " ")
}
