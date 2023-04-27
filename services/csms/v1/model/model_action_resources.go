package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ActionResources struct {

	// 资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	ResourceDetail *Secret `json:"resource_detail,omitempty"`

	// 资源名称，默认为空字符串。
	ResourceName *string `json:"resource_name,omitempty"`

	// 标签列表，没有标签，数组默认为空。
	Tags *[]TagItem `json:"tags,omitempty"`

	// 标签列表，key和value键值对的集合。  - key：表示标签键，一个凭据下最多包含10个key，key不能为空，不能重复，同一个key中value不能重复。key最大长度为36个字符。  - value：表示标签值。每个值最大长度43个字符，value之间为“与”的关系。
	SysTags *[]TagItem `json:"sys_tags,omitempty"`
}

func (o ActionResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionResources struct{}"
	}

	return strings.Join([]string{"ActionResources", string(data)}, " ")
}
