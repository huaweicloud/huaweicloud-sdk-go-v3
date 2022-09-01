package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 搜索字段,key为要匹配的字段，如resource_name等。value为匹配的值。key为固定字典值，不能包含重复的key或不支持的key。 根据key的值确认是否需要模糊匹配，如resource_name默认为模糊搜索（不区分大小写，不支持*，支持字符串匹配），如果value为空字符串则返回空列表（IEF服务不存在资源名称为空的情况，因此这类情况返回空列表）。
type Matches struct {

	// 键。限定为resource_name,后续扩展。
	Key string `json:"key" xml:"key"`

	// 值。每个值最大长度64个unicode字符。不校验字符集范围。
	Value *string `json:"value,omitempty" xml:"value"`
}

func (o Matches) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Matches struct{}"
	}

	return strings.Join([]string{"Matches", string(data)}, " ")
}
