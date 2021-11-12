package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 包含标签，最多包含10个key，每 个key下面的value最多10个，每 个key对应的value可以为空数组但 结构体不能缺失。Key不能重复， 同一个key中values不能重复。结 果返回包含所有标签的资源列表， key之间是与的关系，key-value结 构中value是或的关系。无tag过滤 条件时返回全量数据。
type TagValuesList struct {
	// 键。最大长度127个unicode字 符。key不能为空。(搜索时不对此 参数做字符集校验)，key不能为空 或者空字符串，不能为空格，校验 和使用之前先trim前后半角空格。

	Key string `json:"key"`
	// 值列表。每个值最大长度255个 unicode字符，校验和使用之前先 trim前后半角空格。 value可为空数组但不可缺省。 如果values为空列表，则表示 any_value（查询任意value）。 value之间为或的关系。 (搜索时不对此参数做字符集校 验，只做长度校验)

	Values []string `json:"values"`
}

func (o TagValuesList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagValuesList struct{}"
	}

	return strings.Join([]string{"TagValuesList", string(data)}, " ")
}
