package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DomainMigrateProjectStatus
type DomainMigrateProjectStatus struct {

	// 迁移状态
	Status DomainMigrateProjectStatusStatus `json:"status"`

	// 项目ID
	ProjectId string `json:"project_id"`

	// 项目名称
	ProjectName string `json:"project_name"`

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 迁移进度
	Progress *int32 `json:"progress,omitempty"`

	// 失败错误码（仅当项目状态为失败时才有该参数）。
	FailCode *int32 `json:"fail_code,omitempty"`

	// 失败原因（仅当项目状态为失败时才有该参数）。
	FailReason *string `json:"fail_reason,omitempty"`
}

func (o DomainMigrateProjectStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainMigrateProjectStatus struct{}"
	}

	return strings.Join([]string{"DomainMigrateProjectStatus", string(data)}, " ")
}

type DomainMigrateProjectStatusStatus struct {
	value string
}

type DomainMigrateProjectStatusStatusEnum struct {
	MIGRATING DomainMigrateProjectStatusStatus
	SUCCESS   DomainMigrateProjectStatusStatus
	FAILED    DomainMigrateProjectStatusStatus
}

func GetDomainMigrateProjectStatusStatusEnum() DomainMigrateProjectStatusStatusEnum {
	return DomainMigrateProjectStatusStatusEnum{
		MIGRATING: DomainMigrateProjectStatusStatus{
			value: "migrating",
		},
		SUCCESS: DomainMigrateProjectStatusStatus{
			value: "success",
		},
		FAILED: DomainMigrateProjectStatusStatus{
			value: "failed",
		},
	}
}

func (c DomainMigrateProjectStatusStatus) Value() string {
	return c.value
}

func (c DomainMigrateProjectStatusStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DomainMigrateProjectStatusStatus) UnmarshalJSON(b []byte) error {
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
