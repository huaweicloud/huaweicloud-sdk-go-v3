package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountSubNetworkInterfacesByTagsRequestBody This is a auto create Body Object
type CountSubNetworkInterfacesByTagsRequestBody struct {

	// **参数解释**： 包含标签。结果返回包含所有标签的资源列表，无tag过滤条件时返回全量数据。 **约束限制**： - key之间是与的关系，key-value结构中value是或的关系。 - 最多包含50个key，每个key下面的value最多10个。 - 每个key对应的value可以为空数组但结构体不能缺失。 - Key不能重复，同一个key中values不能重复。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Tags *[]ListTag `json:"tags,omitempty"`
}

func (o CountSubNetworkInterfacesByTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountSubNetworkInterfacesByTagsRequestBody struct{}"
	}

	return strings.Join([]string{"CountSubNetworkInterfacesByTagsRequestBody", string(data)}, " ")
}
