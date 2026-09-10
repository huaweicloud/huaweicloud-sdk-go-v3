package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSubNetworkInterfaceTagsRequestBody This is a auto create Body Object
type BatchDeleteSubNetworkInterfaceTagsRequestBody struct {

	// **参数解释**： 标签列表。 **约束限制**： 最大支持20组标签键值对。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Tags []DeleteResourceTagRequestBody `json:"tags"`
}

func (o BatchDeleteSubNetworkInterfaceTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSubNetworkInterfaceTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteSubNetworkInterfaceTagsRequestBody", string(data)}, " ")
}
