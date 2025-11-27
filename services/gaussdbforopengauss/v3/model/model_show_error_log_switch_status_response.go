package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowErrorLogSwitchStatusResponse Response Object
type ShowErrorLogSwitchStatusResponse struct {

	// **参数解释**: 采集状态。 **取值范围**: - ON：开启采集 - OFF：关闭采集
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowErrorLogSwitchStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowErrorLogSwitchStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowErrorLogSwitchStatusResponse", string(data)}, " ")
}
