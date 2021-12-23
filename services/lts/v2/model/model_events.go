package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Events struct {
	// 告警详情

	Annotations *Annotations `json:"annotations"`
	// 告警信息

	Metadata *Metadata `json:"metadata"`
	// 到达时间(时间戳)

	ArrivesAt *interface{} `json:"arrives_at"`
	// 告警清除时间(时间戳)

	EndsAt *interface{} `json:"ends_at"`
	// 告警id

	Id string `json:"id"`
	// 告警产生时间(时间戳)

	StartsAt *interface{} `json:"starts_at"`
	// 告警自动清除时间(时间戳)

	Timeout *interface{} `json:"timeout"`
	// 告警规则类型(SQL/关键词)

	Type string `json:"type"`
}

func (o Events) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Events struct{}"
	}

	return strings.Join([]string{"Events", string(data)}, " ")
}
