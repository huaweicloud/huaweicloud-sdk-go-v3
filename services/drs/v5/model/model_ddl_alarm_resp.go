package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DdlAlarmResp DDL告警信息响应体。
type DdlAlarmResp struct {

	// 记录唯一序号。
	Seqno int64 `json:"seqno"`

	// 数据源库位点。
	Checkpoint string `json:"checkpoint"`

	// DDL告警状态。0无告警，1告警中。
	Status int32 `json:"status"`

	// DDL在源库执行时间。
	DdlTimestamp int64 `json:"ddl_timestamp"`

	// DDL内容。
	DdlText string `json:"ddl_text"`

	// DDL执行结果。false执行失败，true执行成功。
	ExeResult bool `json:"exe_result"`

	// 数据记录时间。
	RecordTime int64 `json:"record_time"`

	// DDL告警清理时间。
	CleanTime *int64 `json:"clean_time,omitempty"`
}

func (o DdlAlarmResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DdlAlarmResp struct{}"
	}

	return strings.Join([]string{"DdlAlarmResp", string(data)}, " ")
}
