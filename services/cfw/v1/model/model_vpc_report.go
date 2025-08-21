package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VpcReport struct {

	// **参数解释**： TOP应用数量 **取值范围**： 不涉及
	App *[]ItemVo `json:"app,omitempty"`

	// **参数解释**： TOP访问目的IP **取值范围**： 不涉及
	DstIp *[]ItemVo `json:"dst_ip,omitempty"`

	Overview *Overview `json:"overview,omitempty"`

	// **参数解释**： TOP访问源IP **取值范围**： 不涉及
	SrcIp *[]ItemVo `json:"src_ip,omitempty"`

	// **参数解释**： 流量趋势 **取值范围**： 不涉及
	TrafficTrend *[]TrendVo `json:"traffic_trend,omitempty"`

	Vpc *Vpc `json:"vpc,omitempty"`
}

func (o VpcReport) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcReport struct{}"
	}

	return strings.Join([]string{"VpcReport", string(data)}, " ")
}
