package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagFilterRequestBody 标签过滤集群请求体。
type TagFilterRequestBody struct {

	// **参数解释**： 过滤的标签对象。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Tags *[]TagMultiValue `json:"tags,omitempty"`
}

func (o TagFilterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagFilterRequestBody struct{}"
	}

	return strings.Join([]string{"TagFilterRequestBody", string(data)}, " ")
}
