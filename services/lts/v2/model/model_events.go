package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Events struct {
	Annotations *Annotations `json:"annotations" xml:"annotations"`

	Metadata *Metadata `json:"metadata" xml:"metadata"`

	// 到达时间(时间戳)
	ArrivesAt int64 `json:"arrives_at" xml:"arrives_at"`

	// 告警清除时间(时间戳)
	EndsAt int64 `json:"ends_at" xml:"ends_at"`

	// 告警id
	Id string `json:"id" xml:"id"`

	// 告警产生时间(时间戳)
	StartsAt int64 `json:"starts_at" xml:"starts_at"`

	// 告警自动清除时间(时间戳)
	Timeout int64 `json:"timeout" xml:"timeout"`

	// 告警规则类型(SQL/关键词)
	Type string `json:"type" xml:"type"`
}

func (o Events) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Events struct{}"
	}

	return strings.Join([]string{"Events", string(data)}, " ")
}
