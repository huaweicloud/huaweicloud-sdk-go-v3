package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type ScanJob struct {
	// 任务创建时间

	CreateTime *int64 `json:"create_time,omitempty"`
	// 任务执行方式

	Cycle *ScanJobCycle `json:"cycle,omitempty"`
	// 任务ID

	Id *string `json:"id,omitempty"`
	// 任务上一次执行时间

	LastRunTime *int64 `json:"last_run_time,omitempty"`
	// 任务上一次扫描风险等级结果

	LastScanRisk *string `json:"last_scan_risk,omitempty"`
	// 任务名称

	Name *string `json:"name,omitempty"`
	// 任务开启状态

	Open *bool `json:"open,omitempty"`
	// 任务使用的规则组

	RuleGroups *[]string `json:"rule_groups,omitempty"`
	// 任务当前状态

	Status *ScanJobStatus `json:"status,omitempty"`
	// SMN服务通知主题

	TopicUrn *string `json:"topic_urn,omitempty"`
	// 是否使用了NLP进行扫描

	UseNlp *bool `json:"use_nlp,omitempty"`
	// 任务启动时间

	StartTime *int64 `json:"start_time,omitempty"`
}

func (o ScanJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScanJob struct{}"
	}

	return strings.Join([]string{"ScanJob", string(data)}, " ")
}

type ScanJobCycle struct {
	value string
}

type ScanJobCycleEnum struct {
	ONCE  ScanJobCycle
	DAY   ScanJobCycle
	WEEK  ScanJobCycle
	MONTH ScanJobCycle
}

func GetScanJobCycleEnum() ScanJobCycleEnum {
	return ScanJobCycleEnum{
		ONCE: ScanJobCycle{
			value: "ONCE",
		},
		DAY: ScanJobCycle{
			value: "DAY",
		},
		WEEK: ScanJobCycle{
			value: "WEEK",
		},
		MONTH: ScanJobCycle{
			value: "MONTH",
		},
	}
}

func (c ScanJobCycle) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScanJobCycle) UnmarshalJSON(b []byte) error {
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

type ScanJobStatus struct {
	value string
}

type ScanJobStatusEnum struct {
	INIT       ScanJobStatus
	WAITING    ScanJobStatus
	RUNNING    ScanJobStatus
	FAILED     ScanJobStatus
	STOPPED    ScanJobStatus
	FINISHED   ScanJobStatus
	TERMINATED ScanJobStatus
}

func GetScanJobStatusEnum() ScanJobStatusEnum {
	return ScanJobStatusEnum{
		INIT: ScanJobStatus{
			value: "INIT",
		},
		WAITING: ScanJobStatus{
			value: "WAITING",
		},
		RUNNING: ScanJobStatus{
			value: "RUNNING",
		},
		FAILED: ScanJobStatus{
			value: "FAILED",
		},
		STOPPED: ScanJobStatus{
			value: "STOPPED",
		},
		FINISHED: ScanJobStatus{
			value: "FINISHED",
		},
		TERMINATED: ScanJobStatus{
			value: "TERMINATED",
		},
	}
}

func (c ScanJobStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScanJobStatus) UnmarshalJSON(b []byte) error {
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
