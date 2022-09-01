package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 返回的日志组信息
type LogGroup struct {

	// 创建时间
	CreationTime int64 `json:"creation_time" xml:"creation_time"`

	// 日志组名称
	LogGroupName string `json:"log_group_name" xml:"log_group_name"`

	// 日志组ID
	LogGroupId string `json:"log_group_id" xml:"log_group_id"`

	// 日志存储时间 天
	TtlInDays int32 `json:"ttl_in_days" xml:"ttl_in_days"`

	// 日志流所属标签
	Tag map[string]string `json:"tag,omitempty" xml:"tag"`
}

func (o LogGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogGroup struct{}"
	}

	return strings.Join([]string{"LogGroup", string(data)}, " ")
}
