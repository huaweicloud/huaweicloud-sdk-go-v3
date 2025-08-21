package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InternetReport struct {
	Eip *Eip `json:"eip,omitempty"`

	In2out *In2Out `json:"in2out,omitempty"`

	Out2in *Out2in `json:"out2in,omitempty"`

	Overview *Overview `json:"overview,omitempty"`

	// **参数解释**： 流量趋势 **取值范围**： 不涉及
	TrafficTrend *[]TrendVo `json:"traffic_trend,omitempty"`
}

func (o InternetReport) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InternetReport struct{}"
	}

	return strings.Join([]string{"InternetReport", string(data)}, " ")
}
