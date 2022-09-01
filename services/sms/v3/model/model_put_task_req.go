package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 查询指定迁移任务的返回体
type PutTaskReq struct {

	// 任务名称（用户自定义）
	Name *string `json:"name,omitempty" xml:"name"`

	// 任务类型，创建时必选，更新时可选
	Type *PutTaskReqType `json:"type,omitempty" xml:"type"`

	// 操作系统类型，分为WINDOWS和LINUX，创建时必选，更新时可选
	OsType *PutTaskReqOsType `json:"os_type,omitempty" xml:"os_type"`

	// 迁移任务id
	Id *string `json:"id,omitempty" xml:"id"`

	// 进程优先级  0：低  1：标准（默认）  2：高
	Priority *PutTaskReqPriority `json:"priority,omitempty" xml:"priority"`

	// 目的端服务器的区域ID
	RegionId *string `json:"region_id,omitempty" xml:"region_id"`

	// 迁移完成后是否启动目的端服务器  true：启动  false：停止
	StartTargetServer *bool `json:"start_target_server,omitempty" xml:"start_target_server"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	// 目的端服务器的IP地址。  公网迁移时请填写弹性IP地址  专线迁移时请填写私有IP地址
	MigrationIp *string `json:"migration_ip,omitempty" xml:"migration_ip"`

	// 目的端服务器的区域名称
	RegionName *string `json:"region_name,omitempty" xml:"region_name"`

	// 目的端服务器所在项目名称
	ProjectName *string `json:"project_name,omitempty" xml:"project_name"`

	// 目的端服务器所在项目ID
	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`

	// 模板ID
	VmTemplateId *string `json:"vm_template_id,omitempty" xml:"vm_template_id"`

	SourceServer *PostSourceServerBody `json:"source_server,omitempty" xml:"source_server"`

	TargetServer *TargetServer `json:"target_server,omitempty" xml:"target_server"`

	// 任务状态
	State *string `json:"state,omitempty" xml:"state"`

	// 预估完成时间
	EstimateCompleteTime *int64 `json:"estimate_complete_time,omitempty" xml:"estimate_complete_time"`

	// 连接状态
	Connected *bool `json:"connected,omitempty" xml:"connected"`

	// 任务创建时间
	CreateDate *int64 `json:"create_date,omitempty" xml:"create_date"`

	// 任务开始时间
	StartDate *int64 `json:"start_date,omitempty" xml:"start_date"`

	// 任务结束时间
	FinishDate *int64 `json:"finish_date,omitempty" xml:"finish_date"`

	// 迁移速率，单位：MB/S
	MigrateSpeed *float64 `json:"migrate_speed,omitempty" xml:"migrate_speed"`

	// 错误信息
	ErrorJson *string `json:"error_json,omitempty" xml:"error_json"`

	// 任务总耗时
	TotalTime *int64 `json:"total_time,omitempty" xml:"total_time"`

	// 暂时保留float,兼容现网老版本的SMS-Agent
	FloatIp *string `json:"float_ip,omitempty" xml:"float_ip"`

	// 迁移剩余时间（秒）
	RemainSeconds *int64 `json:"remain_seconds,omitempty" xml:"remain_seconds"`

	// 目的端的快照id
	TargetSnapshotId *string `json:"target_snapshot_id,omitempty" xml:"target_snapshot_id"`

	CloneServer *CloneServer `json:"clone_server,omitempty" xml:"clone_server"`

	// 任务包含的子任务列表
	SubTasks *[]SubTask `json:"sub_tasks,omitempty" xml:"sub_tasks"`
}

func (o PutTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutTaskReq struct{}"
	}

	return strings.Join([]string{"PutTaskReq", string(data)}, " ")
}

type PutTaskReqType struct {
	value string
}

type PutTaskReqTypeEnum struct {
	MIGRATE_FILE  PutTaskReqType
	MIGRATE_BLOCK PutTaskReqType
}

func GetPutTaskReqTypeEnum() PutTaskReqTypeEnum {
	return PutTaskReqTypeEnum{
		MIGRATE_FILE: PutTaskReqType{
			value: "MIGRATE_FILE",
		},
		MIGRATE_BLOCK: PutTaskReqType{
			value: "MIGRATE_BLOCK",
		},
	}
}

func (c PutTaskReqType) Value() string {
	return c.value
}

func (c PutTaskReqType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PutTaskReqType) UnmarshalJSON(b []byte) error {
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

type PutTaskReqOsType struct {
	value string
}

type PutTaskReqOsTypeEnum struct {
	WINDOWS PutTaskReqOsType
	LINUX   PutTaskReqOsType
}

func GetPutTaskReqOsTypeEnum() PutTaskReqOsTypeEnum {
	return PutTaskReqOsTypeEnum{
		WINDOWS: PutTaskReqOsType{
			value: "WINDOWS",
		},
		LINUX: PutTaskReqOsType{
			value: "LINUX",
		},
	}
}

func (c PutTaskReqOsType) Value() string {
	return c.value
}

func (c PutTaskReqOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PutTaskReqOsType) UnmarshalJSON(b []byte) error {
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

type PutTaskReqPriority struct {
	value int32
}

type PutTaskReqPriorityEnum struct {
	E_0 PutTaskReqPriority
	E_1 PutTaskReqPriority
	E_2 PutTaskReqPriority
}

func GetPutTaskReqPriorityEnum() PutTaskReqPriorityEnum {
	return PutTaskReqPriorityEnum{
		E_0: PutTaskReqPriority{
			value: 0,
		}, E_1: PutTaskReqPriority{
			value: 1,
		}, E_2: PutTaskReqPriority{
			value: 2,
		},
	}
}

func (c PutTaskReqPriority) Value() int32 {
	return c.value
}

func (c PutTaskReqPriority) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PutTaskReqPriority) UnmarshalJSON(b []byte) error {
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
