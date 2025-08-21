package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlowTrendRespData **参数解释**： 会话趋势统计 **取值范围**： 不涉及
type ShowFlowTrendRespData struct {

	// **参数解释**： 流量趋势数据 **取值范围**： 不涉及
	Records *[]FlowTrendVo `json:"records,omitempty"`

	// **参数解释**： 总数 **取值范围**： 不涉及
	Total *int64 `json:"total,omitempty"`
}

func (o ShowFlowTrendRespData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlowTrendRespData struct{}"
	}

	return strings.Join([]string{"ShowFlowTrendRespData", string(data)}, " ")
}
