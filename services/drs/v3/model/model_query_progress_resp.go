package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 获取指定任务迁移进度响应体
type QueryProgressResp struct {

	// 任务Id
	JobId *string `json:"job_id,omitempty" xml:"job_id"`

	// 迁移百分比
	Progress *string `json:"progress,omitempty" xml:"progress"`

	// 增量迁移时延
	IncreTransDelay *string `json:"incre_trans_delay,omitempty" xml:"incre_trans_delay"`

	// 迁移模式
	TaskMode *QueryProgressRespTaskMode `json:"task_mode,omitempty" xml:"task_mode"`

	// 任务状态
	TransferStatus *string `json:"transfer_status,omitempty" xml:"transfer_status"`

	// 迁移时间，时间戳
	ProcessTime *string `json:"process_time,omitempty" xml:"process_time"`

	// 预计剩余时间
	RemainingTime *string `json:"remaining_time,omitempty" xml:"remaining_time"`

	// 数据，结构，索引迁移进度信息体
	ProgressMap map[string]ProgressInfo `json:"progress_map,omitempty" xml:"progress_map"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg"`
}

func (o QueryProgressResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryProgressResp struct{}"
	}

	return strings.Join([]string{"QueryProgressResp", string(data)}, " ")
}

type QueryProgressRespTaskMode struct {
	value string
}

type QueryProgressRespTaskModeEnum struct {
	FULL_TRANS      QueryProgressRespTaskMode
	INCR_TRANS      QueryProgressRespTaskMode
	FULL_INCR_TRANS QueryProgressRespTaskMode
}

func GetQueryProgressRespTaskModeEnum() QueryProgressRespTaskModeEnum {
	return QueryProgressRespTaskModeEnum{
		FULL_TRANS: QueryProgressRespTaskMode{
			value: "FULL_TRANS: 全量",
		},
		INCR_TRANS: QueryProgressRespTaskMode{
			value: "INCR_TRANS: 增量",
		},
		FULL_INCR_TRANS: QueryProgressRespTaskMode{
			value: "FULL_INCR_TRANS: 全量+增量",
		},
	}
}

func (c QueryProgressRespTaskMode) Value() string {
	return c.value
}

func (c QueryProgressRespTaskMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryProgressRespTaskMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
