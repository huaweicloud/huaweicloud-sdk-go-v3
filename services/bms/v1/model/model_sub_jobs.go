package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SubJobs sub_jobs字段数据结构说明
type SubJobs struct {

	// Job的状态。SUCCESS：成功RUNNING：运行中FAIL：失败INIT：正在初始化
	Status *SubJobsStatus `json:"status,omitempty"`

	Entities *Entitie `json:"entities,omitempty"`

	// Job ID
	JobId *string `json:"job_id,omitempty"`

	// Job的类型，包含以下类型：baremetalSingleCreate：创建单个裸金属服务器；baremetalSingleOperate：修改单个裸金属服务器电源状态；baremetalAttachSingleVolume：挂载单个共享磁盘
	JobType *string `json:"job_type,omitempty"`

	// 开始时间。时间戳格式为ISO 8601，例如：2019-04-25T20:04:47.591Z
	BeginTime *string `json:"begin_time,omitempty"`

	// 结束时间。时间戳格式为ISO 8601，例如：2019-04-26T20:04:47.591Z
	EndTime *string `json:"end_time,omitempty"`

	// Job执行失败时的错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// Job执行失败时的错误原因
	FailReason *string `json:"fail_reason,omitempty"`

	// 出现错误时，返回的错误消息
	Message *string `json:"message,omitempty"`

	// 出现错误时，返回的错误码
	Code *string `json:"code,omitempty"`
}

func (o SubJobs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubJobs struct{}"
	}

	return strings.Join([]string{"SubJobs", string(data)}, " ")
}

type SubJobsStatus struct {
	value string
}

type SubJobsStatusEnum struct {
	SUCCESS SubJobsStatus
	RUNNING SubJobsStatus
	FAIL    SubJobsStatus
	INIT    SubJobsStatus
}

func GetSubJobsStatusEnum() SubJobsStatusEnum {
	return SubJobsStatusEnum{
		SUCCESS: SubJobsStatus{
			value: "SUCCESS",
		},
		RUNNING: SubJobsStatus{
			value: "RUNNING",
		},
		FAIL: SubJobsStatus{
			value: "FAIL",
		},
		INIT: SubJobsStatus{
			value: "INIT",
		},
	}
}

func (c SubJobsStatus) Value() string {
	return c.value
}

func (c SubJobsStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SubJobsStatus) UnmarshalJSON(b []byte) error {
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
