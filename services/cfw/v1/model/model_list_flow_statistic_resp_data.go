package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFlowStatisticRespData **参数解释**： 流量日志统计 **取值范围**： 不涉及
type ListFlowStatisticRespData struct {

	// **参数解释**： 统计记录 **取值范围**： 不涉及
	Records *[]FlowDetailsVo `json:"records,omitempty"`

	// **参数解释**： 总数 **取值范围**： 不涉及
	Total *int64 `json:"total,omitempty"`
}

func (o ListFlowStatisticRespData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlowStatisticRespData struct{}"
	}

	return strings.Join([]string{"ListFlowStatisticRespData", string(data)}, " ")
}
