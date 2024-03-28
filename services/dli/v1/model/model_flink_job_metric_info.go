package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkJobMetricInfo 所有输入流。
type FlinkJobMetricInfo struct {

	// 输入流或输出流名称。
	Name *string `json:"name,omitempty"`

	// 总记录数。
	Records *int64 `json:"records,omitempty"`

	// 脏数据记录数。
	CorruptedRecords *int64 `json:"corrupted_records,omitempty"`
}

func (o FlinkJobMetricInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkJobMetricInfo struct{}"
	}

	return strings.Join([]string{"FlinkJobMetricInfo", string(data)}, " ")
}
