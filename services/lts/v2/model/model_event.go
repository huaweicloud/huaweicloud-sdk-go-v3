package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Event struct {
	// 告警信息

	Metadata *Metadata `json:"metadata"`
	// 告警产生时间(时间戳)

	StartsAt *interface{} `json:"starts_at"`
}

func (o Event) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Event struct{}"
	}

	return strings.Join([]string{"Event", string(data)}, " ")
}
