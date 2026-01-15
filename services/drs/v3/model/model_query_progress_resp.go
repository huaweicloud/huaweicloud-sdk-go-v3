package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// QueryProgressResp 获取指定任务迁移进度响应体
type QueryProgressResp struct {

	// 任务Id
	JobId *string `json:"job_id,omitempty"`

	// 迁移百分比
	Progress *string `json:"progress,omitempty"`

	// 增量迁移时延。单位：s
	IncreTransDelay *string `json:"incre_trans_delay,omitempty"`

	// 增量迁移时延。单位：ms
	IncreTransDelayMillis *string `json:"incre_trans_delay_millis,omitempty"`

	// 迁移模式。 - FULL_TRANS: 全量 - INCR_TRANS: 增量 - FULL_INCR_TRANS: 全量+增量
	TaskMode *QueryProgressRespTaskMode `json:"task_mode,omitempty"`

	// 任务状态 CREATING：创建中，CREATE_FAILED：创建失败，CONFIGURATION：配置中，WAITING_FOR_START：等待启动中，RELEASE_RESOURCE_COMPLETE：已结束，DELETED：已删除，INCRE_TRANSFER_STARTED：增量迁移中，INCRE_TRANSFER_FAILED：增量迁移失败，FULL_TRANSFER_STARTED：全量迁移中，FULL_TRANSFER_COMPLETE：全量迁移完成，PAUSING：暂停中，FULL_TRANSFER_FAILED：全量迁移失败
	TransferStatus *string `json:"transfer_status,omitempty"`

	// 迁移时间，时间戳
	ProcessTime *string `json:"process_time,omitempty"`

	// 预计剩余时间
	RemainingTime *string `json:"remaining_time,omitempty"`

	// 数据，结构，索引迁移进度信息体
	ProgressMap map[string]ProgressInfo `json:"progress_map,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`
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
	FULL_INCR_TRANS QueryProgressRespTaskMode
	INCR_TRANS      QueryProgressRespTaskMode
}

func GetQueryProgressRespTaskModeEnum() QueryProgressRespTaskModeEnum {
	return QueryProgressRespTaskModeEnum{
		FULL_TRANS: QueryProgressRespTaskMode{
			value: "FULL_TRANS",
		},
		FULL_INCR_TRANS: QueryProgressRespTaskMode{
			value: "FULL_INCR_TRANS",
		},
		INCR_TRANS: QueryProgressRespTaskMode{
			value: "INCR_TRANS",
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
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
