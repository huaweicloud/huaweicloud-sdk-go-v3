package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GetLogBackupReq struct {

	// 节点名称。通过[查询集群详情](ShowClusterDetail.xml)获取instances中的name属性。
	InstanceName string `json:"instance_name"`

	// 日志级别。可查询的日志级别为：INFO，ERROR，DEBUG，WARN。
	Level string `json:"level"`

	// 日志类型。可查询的日志类型为：deprecation，indexingSlow，searchSlow， instance。
	LogType string `json:"log_type"`

	// 指定返回日志的条数，默认返回100条，最大返回10000条日志，且日志大小不超过1MB。
	Limit *int32 `json:"limit,omitempty"`

	// 返回指定时间之前的日志。
	TimeIndex *string `json:"time_index,omitempty"`

	// 基于日志内容字段值需要过滤的关键字，注意搜索到的日志包含关键字。
	Keyword *string `json:"keyword,omitempty"`
}

func (o GetLogBackupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetLogBackupReq struct{}"
	}

	return strings.Join([]string{"GetLogBackupReq", string(data)}, " ")
}
