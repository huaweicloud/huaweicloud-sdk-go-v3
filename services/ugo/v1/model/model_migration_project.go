package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type MigrationProject struct {

	// 迁移项目ID。
	MigrationProjectId int32 `json:"migration_project_id"`

	// 迁移项目名称。
	MigrationProjectName string `json:"migration_project_name"`

	// 评估项目ID。
	EvaluationProjectId int32 `json:"evaluation_project_id"`

	// 评估项目名称。
	EvaluationProjectName string `json:"evaluation_project_name"`

	// 迁移项目状态。
	MigrationProjectStatus MigrationProjectMigrationProjectStatus `json:"migration_project_status"`

	// 目标库权限检查状态。
	PermissionCheckStatus MigrationProjectPermissionCheckStatus `json:"permission_check_status"`

	// 资源ID。
	ResourceId string `json:"resource_id"`

	// 创建时间。
	CreatedTime string `json:"created_time"`

	// 更新时间。
	UpdatedTime string `json:"updated_time"`
}

func (o MigrationProject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrationProject struct{}"
	}

	return strings.Join([]string{"MigrationProject", string(data)}, " ")
}

type MigrationProjectMigrationProjectStatus struct {
	value string
}

type MigrationProjectMigrationProjectStatusEnum struct {
	READY     MigrationProjectMigrationProjectStatus
	NOT_READY MigrationProjectMigrationProjectStatus
}

func GetMigrationProjectMigrationProjectStatusEnum() MigrationProjectMigrationProjectStatusEnum {
	return MigrationProjectMigrationProjectStatusEnum{
		READY: MigrationProjectMigrationProjectStatus{
			value: "READY",
		},
		NOT_READY: MigrationProjectMigrationProjectStatus{
			value: "NOT_READY",
		},
	}
}

func (c MigrationProjectMigrationProjectStatus) Value() string {
	return c.value
}

func (c MigrationProjectMigrationProjectStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MigrationProjectMigrationProjectStatus) UnmarshalJSON(b []byte) error {
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

type MigrationProjectPermissionCheckStatus struct {
	value string
}

type MigrationProjectPermissionCheckStatusEnum struct {
	SUCCESS MigrationProjectPermissionCheckStatus
	FAILED  MigrationProjectPermissionCheckStatus
	WAITING MigrationProjectPermissionCheckStatus
	PENDING MigrationProjectPermissionCheckStatus
	IGNORE  MigrationProjectPermissionCheckStatus
}

func GetMigrationProjectPermissionCheckStatusEnum() MigrationProjectPermissionCheckStatusEnum {
	return MigrationProjectPermissionCheckStatusEnum{
		SUCCESS: MigrationProjectPermissionCheckStatus{
			value: "SUCCESS",
		},
		FAILED: MigrationProjectPermissionCheckStatus{
			value: "FAILED",
		},
		WAITING: MigrationProjectPermissionCheckStatus{
			value: "WAITING",
		},
		PENDING: MigrationProjectPermissionCheckStatus{
			value: "PENDING",
		},
		IGNORE: MigrationProjectPermissionCheckStatus{
			value: "IGNORE",
		},
	}
}

func (c MigrationProjectPermissionCheckStatus) Value() string {
	return c.value
}

func (c MigrationProjectPermissionCheckStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MigrationProjectPermissionCheckStatus) UnmarshalJSON(b []byte) error {
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
