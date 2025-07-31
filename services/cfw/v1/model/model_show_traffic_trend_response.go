package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTrafficTrendResponse Response Object
type ShowTrafficTrendResponse struct {

	// **参数解释**： 流量趋势数据 **取值范围**： 不涉及
	Data           *[]TrafficTrendVo `json:"data,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowTrafficTrendResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTrafficTrendResponse struct{}"
	}

	return strings.Join([]string{"ShowTrafficTrendResponse", string(data)}, " ")
}
