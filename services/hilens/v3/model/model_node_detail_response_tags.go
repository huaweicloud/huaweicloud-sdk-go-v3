package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeDetailResponseTags 设备标签对列表
type NodeDetailResponseTags struct {

	// 标签键，最大长度36个字符。不能为空，只能包含大小写字母，数字，中划线“-”，下划线“_”
	Key *string `json:"key,omitempty"`

	// 标签值，每个值最大长度43个字符，删除时如果value有值按照key/value删除，如果value没值，则按照key删除。不能为空，只能包含大小写字母，数字，中划线“-”，下划线“_”
	Value *string `json:"value,omitempty"`
}

func (o NodeDetailResponseTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeDetailResponseTags struct{}"
	}

	return strings.Join([]string{"NodeDetailResponseTags", string(data)}, " ")
}
