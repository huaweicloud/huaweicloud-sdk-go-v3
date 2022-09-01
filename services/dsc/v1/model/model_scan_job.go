package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ScanJob struct {

	// 任务创建时间
	CreateTime *int64 `json:"create_time,omitempty" xml:"create_time"`

	// 任务执行方式
	Cycle *ScanJobCycle `json:"cycle,omitempty" xml:"cycle"`

	// 任务ID
	Id *string `json:"id,omitempty" xml:"id"`

	// 任务上一次执行时间
	LastRunTime *int64 `json:"last_run_time,omitempty" xml:"last_run_time"`

	// 任务上一次扫描风险等级结果
	LastScanRisk *string `json:"last_scan_risk,omitempty" xml:"last_scan_risk"`

	// 任务名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 任务开启状态
	Open *bool `json:"open,omitempty" xml:"open"`

	// 任务使用的规则组
	RuleGroups *[]string `json:"rule_groups,omitempty" xml:"rule_groups"`

	// 任务当前状态
	Status *ScanJobStatus `json:"status,omitempty" xml:"status"`

	// SMN服务通知主题
	TopicUrn *string `json:"topic_urn,omitempty" xml:"topic_urn"`

	// 是否使用了NLP进行扫描
	UseNlp *bool `json:"use_nlp,omitempty" xml:"use_nlp"`

	// 任务启动时间
	StartTime *int64 `json:"start_time,omitempty" xml:"start_time"`
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

func (c ScanJobCycle) Value() string {
	return c.value
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

func (c ScanJobStatus) Value() string {
	return c.value
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
