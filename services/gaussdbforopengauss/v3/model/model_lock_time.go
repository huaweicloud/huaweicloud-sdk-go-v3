package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LockTime struct {

	// **参数解释**: 总耗时（单位：微秒）。 **取值范围**: 不涉及。
	AllTime int64 `json:"all_time"`

	LockTimeDetails *EventTimeInfo `json:"lock_time_details"`
}

func (o LockTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LockTime struct{}"
	}

	return strings.Join([]string{"LockTime", string(data)}, " ")
}
