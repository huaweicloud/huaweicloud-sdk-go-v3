package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// OfflineKeyAnalysisRecord 单条离线全量key分析执行记录
type OfflineKeyAnalysisRecord struct {

	// **参数解释**： 任务执行记录ID（此ID对应“查询离线全量key分析详情”参数中的任务ID）。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 离线分析任务状态。 **取值范围**： waiting：任务等待中。 running：任务进行中。 success：任务执行成功。 failed：任务执行失败。
	Status *OfflineKeyAnalysisRecordStatus `json:"status,omitempty"`

	// **参数解释**： 分析任务创建时间。 **取值范围**： 不涉及。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释**： 分析任务开始时间。 **取值范围**： 不涉及。
	StartedAt *string `json:"started_at,omitempty"`

	// **参数解释**： 分析任务结束时间。 **取值范围**： 不涉及。
	FinishedAt *string `json:"finished_at,omitempty"`
}

func (o OfflineKeyAnalysisRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OfflineKeyAnalysisRecord struct{}"
	}

	return strings.Join([]string{"OfflineKeyAnalysisRecord", string(data)}, " ")
}

type OfflineKeyAnalysisRecordStatus struct {
	value string
}

type OfflineKeyAnalysisRecordStatusEnum struct {
	WAITING OfflineKeyAnalysisRecordStatus
	RUNNING OfflineKeyAnalysisRecordStatus
	SUCCESS OfflineKeyAnalysisRecordStatus
	FAILED  OfflineKeyAnalysisRecordStatus
}

func GetOfflineKeyAnalysisRecordStatusEnum() OfflineKeyAnalysisRecordStatusEnum {
	return OfflineKeyAnalysisRecordStatusEnum{
		WAITING: OfflineKeyAnalysisRecordStatus{
			value: "waiting",
		},
		RUNNING: OfflineKeyAnalysisRecordStatus{
			value: "running",
		},
		SUCCESS: OfflineKeyAnalysisRecordStatus{
			value: "success",
		},
		FAILED: OfflineKeyAnalysisRecordStatus{
			value: "failed",
		},
	}
}

func (c OfflineKeyAnalysisRecordStatus) Value() string {
	return c.value
}

func (c OfflineKeyAnalysisRecordStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OfflineKeyAnalysisRecordStatus) UnmarshalJSON(b []byte) error {
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
