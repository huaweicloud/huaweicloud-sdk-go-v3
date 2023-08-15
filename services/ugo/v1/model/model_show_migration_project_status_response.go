package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowMigrationProjectStatusResponse Response Object
type ShowMigrationProjectStatusResponse struct {

	// 迁移项目ID。
	MigrationProjectId *int32 `json:"migration_project_id,omitempty"`

	// 迁移项目名称。
	MigrationProjectName *string `json:"migration_project_name,omitempty"`

	// 评估项目ID。
	EvaluationProjectId *int32 `json:"evaluation_project_id,omitempty"`

	// 评估项目名称。
	EvaluationProjectName *string `json:"evaluation_project_name,omitempty"`

	// 迁移项目状态。
	MigrationProjectStatus *ShowMigrationProjectStatusResponseMigrationProjectStatus `json:"migration_project_status,omitempty"`

	// 目标库权限检查状态。
	PermissionCheckStatus *ShowMigrationProjectStatusResponsePermissionCheckStatus `json:"permission_check_status,omitempty"`

	// 资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 创建时间。
	CreatedTime *string `json:"created_time,omitempty"`

	// 更新时间。
	UpdatedTime    *string `json:"updated_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowMigrationProjectStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMigrationProjectStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowMigrationProjectStatusResponse", string(data)}, " ")
}

type ShowMigrationProjectStatusResponseMigrationProjectStatus struct {
	value string
}

type ShowMigrationProjectStatusResponseMigrationProjectStatusEnum struct {
	READY     ShowMigrationProjectStatusResponseMigrationProjectStatus
	NOT_READY ShowMigrationProjectStatusResponseMigrationProjectStatus
}

func GetShowMigrationProjectStatusResponseMigrationProjectStatusEnum() ShowMigrationProjectStatusResponseMigrationProjectStatusEnum {
	return ShowMigrationProjectStatusResponseMigrationProjectStatusEnum{
		READY: ShowMigrationProjectStatusResponseMigrationProjectStatus{
			value: "READY",
		},
		NOT_READY: ShowMigrationProjectStatusResponseMigrationProjectStatus{
			value: "NOT_READY",
		},
	}
}

func (c ShowMigrationProjectStatusResponseMigrationProjectStatus) Value() string {
	return c.value
}

func (c ShowMigrationProjectStatusResponseMigrationProjectStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMigrationProjectStatusResponseMigrationProjectStatus) UnmarshalJSON(b []byte) error {
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

type ShowMigrationProjectStatusResponsePermissionCheckStatus struct {
	value string
}

type ShowMigrationProjectStatusResponsePermissionCheckStatusEnum struct {
	SUCCESS ShowMigrationProjectStatusResponsePermissionCheckStatus
	FAILED  ShowMigrationProjectStatusResponsePermissionCheckStatus
	WAITING ShowMigrationProjectStatusResponsePermissionCheckStatus
	PENDING ShowMigrationProjectStatusResponsePermissionCheckStatus
	IGNORE  ShowMigrationProjectStatusResponsePermissionCheckStatus
}

func GetShowMigrationProjectStatusResponsePermissionCheckStatusEnum() ShowMigrationProjectStatusResponsePermissionCheckStatusEnum {
	return ShowMigrationProjectStatusResponsePermissionCheckStatusEnum{
		SUCCESS: ShowMigrationProjectStatusResponsePermissionCheckStatus{
			value: "SUCCESS",
		},
		FAILED: ShowMigrationProjectStatusResponsePermissionCheckStatus{
			value: "FAILED",
		},
		WAITING: ShowMigrationProjectStatusResponsePermissionCheckStatus{
			value: "WAITING",
		},
		PENDING: ShowMigrationProjectStatusResponsePermissionCheckStatus{
			value: "PENDING",
		},
		IGNORE: ShowMigrationProjectStatusResponsePermissionCheckStatus{
			value: "IGNORE",
		},
	}
}

func (c ShowMigrationProjectStatusResponsePermissionCheckStatus) Value() string {
	return c.value
}

func (c ShowMigrationProjectStatusResponsePermissionCheckStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMigrationProjectStatusResponsePermissionCheckStatus) UnmarshalJSON(b []byte) error {
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
