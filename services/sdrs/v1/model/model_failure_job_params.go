package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 失败任务数据结构
type FailureJobParams struct {
	// 任务名称。

	JobType string `json:"job_type"`
	// 任务状态。当前仅支持“FAIL”。FAIL：表示任务失败。

	JobStatus FailureJobParamsJobStatus `json:"job_status"`
	// 任务操作时间。默认格式为：\"yyyy-MM-ddTHH:mm:ss.SSSZ\"，例：\"2019-04-01T12:00:00.000Z\"。

	BeginTime string `json:"begin_time"`
	// 任务id。执行异步API命令下发成功的返回参数。

	JobId string `json:"job_id"`
	// 失败任务状态。createFail：表示创建失败。deleteFail：表示删除失败。attachFail：表示挂载失败。detachFail：表示卸载失败。expandFail：表示扩容失败。resizeFail：表示变更规格失败。startFail：表示开启保护失败。stopFail：表示停止保护失败。reverseFail：表示切换失败。failoverFail：表示故障切换失败。reprotectFail : 表示重保护失败。

	FailureStatus FailureJobParamsFailureStatus `json:"failure_status"`
	// 资源ID。

	ResourceId string `json:"resource_id"`
	// 资源名称。

	ResourceName string `json:"resource_name"`
	// 任务失败错误码。

	ErrorCode string `json:"error_code"`
	// 任务失败原因。

	FailReason string `json:"fail_reason"`
	// 资源类型。 server_groups：表示保护组。 protected_instances：表示保护实例。 replications：表示复制对。 disaster_recovery_drills：表示容灾演练。

	ResourceType FailureJobParamsResourceType `json:"resource_type"`
}

func (o FailureJobParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FailureJobParams struct{}"
	}

	return strings.Join([]string{"FailureJobParams", string(data)}, " ")
}

type FailureJobParamsJobStatus struct {
	value string
}

type FailureJobParamsJobStatusEnum struct {
	FAIL FailureJobParamsJobStatus
}

func GetFailureJobParamsJobStatusEnum() FailureJobParamsJobStatusEnum {
	return FailureJobParamsJobStatusEnum{
		FAIL: FailureJobParamsJobStatus{
			value: "FAIL",
		},
	}
}

func (c FailureJobParamsJobStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FailureJobParamsJobStatus) UnmarshalJSON(b []byte) error {
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

type FailureJobParamsFailureStatus struct {
	value string
}

type FailureJobParamsFailureStatusEnum struct {
	CREATE_FAIL    FailureJobParamsFailureStatus
	DELETE_FAIL    FailureJobParamsFailureStatus
	ATTACH_FAIL    FailureJobParamsFailureStatus
	DETACH_FAIL    FailureJobParamsFailureStatus
	EXPAND_FAIL    FailureJobParamsFailureStatus
	RESIZE_FAIL    FailureJobParamsFailureStatus
	START_FAIL     FailureJobParamsFailureStatus
	STOP_FAIL      FailureJobParamsFailureStatus
	REVERSE_FAIL   FailureJobParamsFailureStatus
	FAILOVER_FAIL  FailureJobParamsFailureStatus
	REPROTECT_FAIL FailureJobParamsFailureStatus
}

func GetFailureJobParamsFailureStatusEnum() FailureJobParamsFailureStatusEnum {
	return FailureJobParamsFailureStatusEnum{
		CREATE_FAIL: FailureJobParamsFailureStatus{
			value: "createFail",
		},
		DELETE_FAIL: FailureJobParamsFailureStatus{
			value: "deleteFail",
		},
		ATTACH_FAIL: FailureJobParamsFailureStatus{
			value: "attachFail",
		},
		DETACH_FAIL: FailureJobParamsFailureStatus{
			value: "detachFail",
		},
		EXPAND_FAIL: FailureJobParamsFailureStatus{
			value: "expandFail",
		},
		RESIZE_FAIL: FailureJobParamsFailureStatus{
			value: "resizeFail",
		},
		START_FAIL: FailureJobParamsFailureStatus{
			value: "startFail",
		},
		STOP_FAIL: FailureJobParamsFailureStatus{
			value: "stopFail",
		},
		REVERSE_FAIL: FailureJobParamsFailureStatus{
			value: "reverseFail",
		},
		FAILOVER_FAIL: FailureJobParamsFailureStatus{
			value: "failoverFail",
		},
		REPROTECT_FAIL: FailureJobParamsFailureStatus{
			value: "reprotectFail",
		},
	}
}

func (c FailureJobParamsFailureStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FailureJobParamsFailureStatus) UnmarshalJSON(b []byte) error {
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

type FailureJobParamsResourceType struct {
	value string
}

type FailureJobParamsResourceTypeEnum struct {
	SERVER_GROUPS            FailureJobParamsResourceType
	PROTECTED_INSTANCES      FailureJobParamsResourceType
	REPLICATIONS             FailureJobParamsResourceType
	DISASTER_RECOVERY_DRILLS FailureJobParamsResourceType
}

func GetFailureJobParamsResourceTypeEnum() FailureJobParamsResourceTypeEnum {
	return FailureJobParamsResourceTypeEnum{
		SERVER_GROUPS: FailureJobParamsResourceType{
			value: "server_groups",
		},
		PROTECTED_INSTANCES: FailureJobParamsResourceType{
			value: "protected_instances",
		},
		REPLICATIONS: FailureJobParamsResourceType{
			value: "replications",
		},
		DISASTER_RECOVERY_DRILLS: FailureJobParamsResourceType{
			value: "disaster_recovery_drills",
		},
	}
}

func (c FailureJobParamsResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FailureJobParamsResourceType) UnmarshalJSON(b []byte) error {
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
