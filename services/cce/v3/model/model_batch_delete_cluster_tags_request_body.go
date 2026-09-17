package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteClusterTagsRequestBody 批量删除指定集群资源标签的请求体
type BatchDeleteClusterTagsRequestBody struct {

	// **参数解释：** 待删除的集群资源标签列表。 **约束限制：** 删除时仅需指定key，value将被忽略。 **取值范围：** 1-20个ResourceDeleteTag对象。 **默认取值：** 不涉及
	Tags []ResourceDeleteTag `json:"tags"`
}

func (o BatchDeleteClusterTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteClusterTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteClusterTagsRequestBody", string(data)}, " ")
}
