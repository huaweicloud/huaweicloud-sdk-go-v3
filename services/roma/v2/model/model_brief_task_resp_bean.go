package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 查询任务列表返回的对象
type BriefTaskRespBean struct {

	// 任务ID
	TaskId *string `json:"task_id,omitempty" xml:"task_id"`

	// 任务名称
	TaskName *string `json:"task_name,omitempty" xml:"task_name"`

	// 任务类型 - REALTIME (实时) - TIMING (定时)
	TaskType *BriefTaskRespBeanTaskType `json:"task_type,omitempty" xml:"task_type"`

	// 任务状态 - 0 (停止/未启动) - 1 (运行中)
	Status *BriefTaskRespBeanStatus `json:"status,omitempty" xml:"status"`

	// 创建时间
	CreatedDate *string `json:"created_date,omitempty" xml:"created_date"`

	// 任务的版本
	Version *string `json:"version,omitempty" xml:"version"`

	// 上次修改时间
	LastModifiedTime *int64 `json:"last_modified_time,omitempty" xml:"last_modified_time"`

	// 任务执行状态  - UNSTARTED (未启动)  - WAITING (等待执行)  - RUNNING (执行中)  - SUCCESS (执行成功)  - CANCELLED (任务取消)  - ERROR (执行异常)
	ExecuteStatus *string `json:"execute_status,omitempty" xml:"execute_status"`

	// 源端数据源所属应用ID
	SourceAppId *string `json:"source_app_id,omitempty" xml:"source_app_id"`

	// 目标端数据源所属应用ID
	TargetAppId *string `json:"target_app_id,omitempty" xml:"target_app_id"`

	// 源端实例ID
	SourceInstanceId *string `json:"source_instance_id,omitempty" xml:"source_instance_id"`

	// 目标端实例ID
	TargetInstanceId *string `json:"target_instance_id,omitempty" xml:"target_instance_id"`

	// 组合任务类型, 可为空
	ExtType *string `json:"ext_type,omitempty" xml:"ext_type"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	// 任务标签
	TaskTag *string `json:"task_tag,omitempty" xml:"task_tag"`
}

func (o BriefTaskRespBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BriefTaskRespBean struct{}"
	}

	return strings.Join([]string{"BriefTaskRespBean", string(data)}, " ")
}

type BriefTaskRespBeanTaskType struct {
	value string
}

type BriefTaskRespBeanTaskTypeEnum struct {
	REALTIME BriefTaskRespBeanTaskType
	TIMING   BriefTaskRespBeanTaskType
}

func GetBriefTaskRespBeanTaskTypeEnum() BriefTaskRespBeanTaskTypeEnum {
	return BriefTaskRespBeanTaskTypeEnum{
		REALTIME: BriefTaskRespBeanTaskType{
			value: "REALTIME",
		},
		TIMING: BriefTaskRespBeanTaskType{
			value: "TIMING",
		},
	}
}

func (c BriefTaskRespBeanTaskType) Value() string {
	return c.value
}

func (c BriefTaskRespBeanTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BriefTaskRespBeanTaskType) UnmarshalJSON(b []byte) error {
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

type BriefTaskRespBeanStatus struct {
	value int32
}

type BriefTaskRespBeanStatusEnum struct {
	E_0 BriefTaskRespBeanStatus
	E_1 BriefTaskRespBeanStatus
}

func GetBriefTaskRespBeanStatusEnum() BriefTaskRespBeanStatusEnum {
	return BriefTaskRespBeanStatusEnum{
		E_0: BriefTaskRespBeanStatus{
			value: 0,
		}, E_1: BriefTaskRespBeanStatus{
			value: 1,
		},
	}
}

func (c BriefTaskRespBeanStatus) Value() int32 {
	return c.value
}

func (c BriefTaskRespBeanStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BriefTaskRespBeanStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
