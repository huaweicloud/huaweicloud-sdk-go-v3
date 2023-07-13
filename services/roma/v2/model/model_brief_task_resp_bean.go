package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BriefTaskRespBean 查询任务列表返回的对象
type BriefTaskRespBean struct {

	// 任务ID
	TaskId *string `json:"task_id,omitempty"`

	// 任务名称
	TaskName *string `json:"task_name,omitempty"`

	// 任务类型 - REALTIME (实时) - TIMING (定时)
	TaskType *BriefTaskRespBeanTaskType `json:"task_type,omitempty"`

	// 任务状态 - 0 (停止/未启动) - 1 (运行中)
	Status *BriefTaskRespBeanStatus `json:"status,omitempty"`

	// 创建时间
	CreatedDate *string `json:"created_date,omitempty"`

	// 任务的版本
	Version *string `json:"version,omitempty"`

	// 上次修改时间
	LastModifiedTime *int64 `json:"last_modified_time,omitempty"`

	// 任务执行状态  - UNSTARTED (未启动)  - WAITING (等待执行)  - RUNNING (执行中)  - SUCCESS (执行成功)  - CANCELLED (任务取消)  - ERROR (执行异常)
	ExecuteStatus *string `json:"execute_status,omitempty"`

	// 源端数据源所属应用ID
	SourceAppId *string `json:"source_app_id,omitempty"`

	// 目标端数据源所属应用ID
	TargetAppId *string `json:"target_app_id,omitempty"`

	// 源端实例ID
	SourceInstanceId *string `json:"source_instance_id,omitempty"`

	// 目标端实例ID
	TargetInstanceId *string `json:"target_instance_id,omitempty"`

	// 组合任务类型, 可为空
	ExtType *string `json:"ext_type,omitempty"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 任务标签
	TaskTag *string `json:"task_tag,omitempty"`
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
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
