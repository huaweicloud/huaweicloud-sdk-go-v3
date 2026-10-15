package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSwitchoverInfoResponse Response Object
type ShowSwitchoverInfoResponse struct {

	// 上次倒换完成时间点
	LastSwitchoverTime *string `json:"last_switchover_time,omitempty"`

	// 上次倒换任务位点信息
	IncrementCheckPoint *string `json:"increment_check_point,omitempty"`

	// 上次倒换目标库倒换时位点
	TargetSafePoint *string `json:"target_safe_point,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o ShowSwitchoverInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSwitchoverInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowSwitchoverInfoResponse", string(data)}, " ")
}
