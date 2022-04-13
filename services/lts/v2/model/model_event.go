package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Event struct {
	Metadata *Metadata `json:"metadata"`
	// 告警产生时间(时间戳)

	StartsAt int64 `json:"starts_at"`
}

func (o Event) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Event struct{}"
	}

	return strings.Join([]string{"Event", string(data)}, " ")
}
