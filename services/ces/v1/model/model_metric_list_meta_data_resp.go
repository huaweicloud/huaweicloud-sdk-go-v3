package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetricListMetaDataResp **参数解释**： 查询结果元数据信息，包括分页信息等
type MetricListMetaDataResp struct {

	// **参数解释**： 当前返回结果条数。 **取值范围**： 不涉及
	Count int32 `json:"count"`

	// **参数解释**： 总条数。 **取值范围**： 不涉及
	Total int32 `json:"total"`

	// **参数解释**： 下一个开始的标记，用于分页。 **取值范围**： 不涉及
	Marker string `json:"marker"`
}

func (o MetricListMetaDataResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricListMetaDataResp struct{}"
	}

	return strings.Join([]string{"MetricListMetaDataResp", string(data)}, " ")
}
