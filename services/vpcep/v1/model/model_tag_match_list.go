package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 搜索字段，key为要匹配的字段， 如resource_name等。value为匹 配的值。key为固定字典值，不能 包含重复的key或不支持的key。 根据key的值确认是否需要模糊匹 配，如resource_name默认为模糊 搜索（不区分大小写），如果 value为空字符串精确匹配（多数 服务不存在资源名称为空的情况， 因此此类情况返回空列表）。 resource_id为精确匹配。第一期 只做resource_name，后续再扩 展。
type TagMatchList struct {
	// 键。第一期限定为 resource_name，后续扩展。  Key 字符集规范 键，unicode字符，不能为空。不能包含下列字符： 非打印字符ASCII(0-31)、“=”、“*”、“<”、“>”、“\\”、“,”、“|”和 “/”。

	Key string `json:"key"`
	// 值。每个值最大长度255个 unicode字符。不校验字符集范。  Value 字符集规范 值，unicode字符。不能包含下列字符： 非打印字符ASCII(0-31)、“=”、“*”、“<”、“>”、“\\”、“,”、“|”和 “/”。

	Value string `json:"value"`
}

func (o TagMatchList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagMatchList struct{}"
	}

	return strings.Join([]string{"TagMatchList", string(data)}, " ")
}
