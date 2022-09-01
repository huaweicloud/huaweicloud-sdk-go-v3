package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Shards struct {

	// 节点Id。
	NodeId string `json:"node_id" xml:"node_id"`

	// 组件Id。最大长度7个字符，不能为null或者空字符串，不能为空格，校验和使用之前会自动过滤掉前后空格。至少包含大写字母（A-Z），小写字母（a-z），数字（0-9），非字母数字字符（限定为_）四类字符中的三类字符。组件id通过查询实例的组件列表接口（ListComponentInfos）获取
	ComponentId string `json:"component_id" xml:"component_id"`
}

func (o Shards) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Shards struct{}"
	}

	return strings.Join([]string{"Shards", string(data)}, " ")
}
