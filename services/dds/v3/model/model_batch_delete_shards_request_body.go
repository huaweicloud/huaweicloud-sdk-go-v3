package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteShardsRequestBody 请求体
type BatchDeleteShardsRequestBody struct {

	// **参数解释：** 组ID。可以调用“查询实例列表和详情”接口获取。 **约束限制：** 只支持shard组。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Ids []string `json:"ids"`
}

func (o BatchDeleteShardsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteShardsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteShardsRequestBody", string(data)}, " ")
}
