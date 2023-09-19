package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShootScriptItem 剧本参数配置。
type ShootScriptItem struct {

	// 剧本序号。
	SequenceNo *int32 `json:"sequence_no,omitempty"`

	// 开始时间。  单位秒。  相对于内容的开始时间。 > 预留字段。当前只需要填sequence_no即可。
	StartTime *float32 `json:"start_time,omitempty"`

	// 结束时间。  单位秒。  相对于内容的结束时间。 > 预留字段。当前只需要填sequence_no即可。
	EndTime *float32 `json:"end_time,omitempty"`

	ShootScript *ShootScript `json:"shoot_script"`
}

func (o ShootScriptItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShootScriptItem struct{}"
	}

	return strings.Join([]string{"ShootScriptItem", string(data)}, " ")
}
