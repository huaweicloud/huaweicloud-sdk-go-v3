package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogtankItem 云日志信息
type LogtankItem struct {

	// 云日志信息唯一的资源标识。
	Id string `json:"id"`

	// 云日志服务日志组ID。
	LogGroupId string `json:"log_group_id"`

	// 云日志服务日志流ID。
	LogStreamId string `json:"log_stream_id"`

	// 创建时间。时间格式为UTC时间，YYYY-MM-DDTHH:MM:SSZ。
	CreateTime string `json:"create_time"`

	// 更新时间。时间格式为UTC时间，YYYY-MM-DDTHH:MM:SSZ。
	UpdateTime string `json:"update_time"`
}

func (o LogtankItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogtankItem struct{}"
	}

	return strings.Join([]string{"LogtankItem", string(data)}, " ")
}
