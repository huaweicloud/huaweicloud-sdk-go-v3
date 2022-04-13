package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 子Job的详细信息。
type SubJob struct {
	// 子job的状态。SUCCESS：成功。RUNNING：运行中。FAIL：失败。INIT：正在初始化。

	Status SubJobStatus `json:"status"`

	Entities *SubJobEntities `json:"entities"`
	// 子job的ID。

	JobId string `json:"job_id"`
	// 子job的类型。createVolume：创建单个云硬盘。batchCreateVolume：批量创建云硬盘。deleteVolume：删除单个云硬盘。extendVolume：扩容云硬盘。bulkDeleteVolume：批量删除云硬盘。deleteSingleVolume：批量删除时逐个删除单个云硬盘。retypeVolume：对云硬盘做硬盘类型变更。

	JobType string `json:"job_type"`
	// 开始时间。

	BeginTime string `json:"begin_time"`
	// 结束时间。

	EndTime string `json:"end_time"`
	// 子job执行失败时的错误码。

	ErrorCode string `json:"error_code"`
	// 子job执行失败时的错误原因。

	FailReason string `json:"fail_reason"`
}

func (o SubJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubJob struct{}"
	}

	return strings.Join([]string{"SubJob", string(data)}, " ")
}

type SubJobStatus struct {
	value string
}

type SubJobStatusEnum struct {
	SUCCESS SubJobStatus
	RUNNING SubJobStatus
	FAIL    SubJobStatus
	INIT    SubJobStatus
}

func GetSubJobStatusEnum() SubJobStatusEnum {
	return SubJobStatusEnum{
		SUCCESS: SubJobStatus{
			value: "SUCCESS",
		},
		RUNNING: SubJobStatus{
			value: "RUNNING",
		},
		FAIL: SubJobStatus{
			value: "FAIL",
		},
		INIT: SubJobStatus{
			value: "INIT",
		},
	}
}

func (c SubJobStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SubJobStatus) UnmarshalJSON(b []byte) error {
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
