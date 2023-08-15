package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Events struct {
	Annotations *Annotations `json:"annotations"`

	Metadata *Metadata `json:"metadata"`

	// 到达时间(时间戳)
	ArrivesAt int64 `json:"arrives_at"`

	// 告警清除时间(时间戳)
	EndsAt int64 `json:"ends_at"`

	// 告警id
	Id string `json:"id"`

	// 告警产生时间(时间戳)
	StartsAt int64 `json:"starts_at"`

	// 告警自动清除时间(时间戳)
	Timeout int64 `json:"timeout"`

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
