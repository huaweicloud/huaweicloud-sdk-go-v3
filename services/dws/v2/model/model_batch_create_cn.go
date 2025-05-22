package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateCn **参数解释**： 批量增加CN节点任务完成后，集群总CN数量。 **约束限制**： 不涉及。 **取值范围**： 大于0，小于剩余dn节点总数。 **默认取值**： 不涉及。
type BatchCreateCn struct {

	// **参数解释**： 批量增加CN节点任务完成后，集群总CN数量。集群支持的CN节点数量与集群当前版本和节点数量相关，具体支持范围可根据“查询集群CN节点”查询，“min_num”为支持的最小数量，max_num为支持的最大数量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Num int32 `json:"num"`
}

func (o BatchCreateCn) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateCn struct{}"
	}

	return strings.Join([]string{"BatchCreateCn", string(data)}, " ")
}
