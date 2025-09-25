package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LwlockTime 轻量级锁耗时信息
type LwlockTime struct {

	// **参数解释**: 总耗时（单位：微秒）。 **取值范围**: 不涉及。
	AllTime int64 `json:"all_time"`

	LwlockTimeDetails *EventTimeInfo `json:"lwlock_time_details"`
}

func (o LwlockTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LwlockTime struct{}"
	}

	return strings.Join([]string{"LwlockTime", string(data)}, " ")
}
