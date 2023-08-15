package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceSpaceInfo 实例空间信息。数据来源于文件系统。已用空间包含数据空间、日志空间和其他空间，其他空间包括引擎产生的临时文件等。
type InstanceSpaceInfo struct {

	// 实例总空间，以字节为单位。GaussDB(for MySQL)不会返回总空间
	TotalSize *int64 `json:"total_size,omitempty"`

	// 已使用空间，以字节为单位
	UsedSize *int64 `json:"used_size,omitempty"`

	// 数据空间，以字节为单位
	DataSize *int64 `json:"data_size,omitempty"`

	// 日志空间，以字节为单位
	LogSize *int64 `json:"log_size,omitempty"`

	// 近七日的数据平均日增长量，以字节为单位
	AvgDailyGrowth *int64 `json:"avg_daily_growth,omitempty"`

	// 最后一次分析的结果时间，毫秒单位时间戳
	LastResultTime *int64 `json:"last_result_time,omitempty"`
}

func (o InstanceSpaceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceSpaceInfo struct{}"
	}

	return strings.Join([]string{"InstanceSpaceInfo", string(data)}, " ")
}
