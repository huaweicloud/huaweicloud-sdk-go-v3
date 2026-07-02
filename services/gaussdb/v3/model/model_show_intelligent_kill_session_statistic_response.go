package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIntelligentKillSessionStatisticResponse Response Object
type ShowIntelligentKillSessionStatisticResponse struct {

	// **参数解释**：  预览智能Kill会话结果列表。
	Statistics     *[]IntelligentKillSessionStatistic `json:"statistics,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o ShowIntelligentKillSessionStatisticResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIntelligentKillSessionStatisticResponse struct{}"
	}

	return strings.Join([]string{"ShowIntelligentKillSessionStatisticResponse", string(data)}, " ")
}
