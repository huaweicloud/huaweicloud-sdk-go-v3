package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SlowLogTopInfo 慢日志Top信息
type SlowLogTopInfo struct {

	// 对象名称
	ObjectName *string `json:"object_name,omitempty"`

	// 数量
	Count *int64 `json:"count,omitempty"`

	// 占比
	Percent *float64 `json:"percent,omitempty"`

	// 总数
	Total *int64 `json:"total,omitempty"`
}

func (o SlowLogTopInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowLogTopInfo struct{}"
	}

	return strings.Join([]string{"SlowLogTopInfo", string(data)}, " ")
}
