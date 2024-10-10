package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// OfflineTaskInfo 备份迁移任务列表信息体。
type OfflineTaskInfo struct {

	// 任务ID。
	Id string `json:"id"`

	// 任务名称。
	Name string `json:"name"`

	// 任务状态。 - TRANSFERRING：恢复中 - SUCCESS：成功 - FAILED：失败 - PRECHECK FAILED：预检查失败
	Status OfflineTaskInfoStatus `json:"status"`

	// 数据库引擎。 - sqlserver：RDS for SQL Server引擎
	EngineType OfflineTaskInfoEngineType `json:"engine_type"`

	// 错误日志。
	ErrorLog *string `json:"error_log,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 任务创建时间。
	CreateTime string `json:"create_time"`

	// 任务完成时间。
	FinishTime *string `json:"finish_time,omitempty"`

	// 企业项目。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o OfflineTaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OfflineTaskInfo struct{}"
	}

	return strings.Join([]string{"OfflineTaskInfo", string(data)}, " ")
}

type OfflineTaskInfoStatus struct {
	value string
}

type OfflineTaskInfoStatusEnum struct {
	TRANSFERRING    OfflineTaskInfoStatus
	SUCCESS         OfflineTaskInfoStatus
	FAILED          OfflineTaskInfoStatus
	PRECHECK_FAILED OfflineTaskInfoStatus
}

func GetOfflineTaskInfoStatusEnum() OfflineTaskInfoStatusEnum {
	return OfflineTaskInfoStatusEnum{
		TRANSFERRING: OfflineTaskInfoStatus{
			value: "TRANSFERRING",
		},
		SUCCESS: OfflineTaskInfoStatus{
			value: "SUCCESS",
		},
		FAILED: OfflineTaskInfoStatus{
			value: "FAILED",
		},
		PRECHECK_FAILED: OfflineTaskInfoStatus{
			value: "PRECHECK FAILED",
		},
	}
}

func (c OfflineTaskInfoStatus) Value() string {
	return c.value
}

func (c OfflineTaskInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OfflineTaskInfoStatus) UnmarshalJSON(b []byte) error {
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

type OfflineTaskInfoEngineType struct {
	value string
}

type OfflineTaskInfoEngineTypeEnum struct {
	SQLSERVER OfflineTaskInfoEngineType
}

func GetOfflineTaskInfoEngineTypeEnum() OfflineTaskInfoEngineTypeEnum {
	return OfflineTaskInfoEngineTypeEnum{
		SQLSERVER: OfflineTaskInfoEngineType{
			value: "sqlserver",
		},
	}
}

func (c OfflineTaskInfoEngineType) Value() string {
	return c.value
}

func (c OfflineTaskInfoEngineType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OfflineTaskInfoEngineType) UnmarshalJSON(b []byte) error {
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
