package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 不包含标签，最多包含10个key， 每个key下面的value最多10个， 每个key对应的value可以为空数组 但结构体不能缺失。Key不能重 复，同一个key中values不能重 复。结果返回不包含标签的资源列 表，key之间是与的关系，keyvalue 结构中value是或的关系。无 过滤条件时返回全量数据。
type NotTagsValuesList struct {
	// 键。最大长度127个unicode字 符。key不能为空。(搜索时不对此 参数做字符集校验)，key不能为空 或者空字符串，不能为空格，校验 和使用之前先trim前后半角空格。

	Key string `json:"key"`
	// 值列表。每个值最大长度255个 unicode字符，校验和使用之前先 trim前后半角空格。 value可为空数组但不可缺省。 如果values为空列表，则表示 any_value（查询任意value）。 value之间为或的关系。 (搜索时不对此参数做字符集校 验，只做长度校验)

	Values []string `json:"values"`
}

func (o NotTagsValuesList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotTagsValuesList struct{}"
	}

	return strings.Join([]string{"NotTagsValuesList", string(data)}, " ")
}
