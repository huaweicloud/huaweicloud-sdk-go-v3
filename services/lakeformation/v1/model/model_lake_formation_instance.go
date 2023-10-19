package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// LakeFormationInstance LakeFormation实例详细信息
type LakeFormationInstance struct {

	// LakeFormation实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称。
	Name *string `json:"name,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 逻辑多租和物理多租的判断。false为物理多租；true为逻辑多租。
	Shared *bool `json:"shared,omitempty"`

	// 是否为默认实例
	DefaultInstance *bool `json:"default_instance,omitempty"`

	// 实例创建时间戳
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 实例更新时间戳
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 实例状态,RESOURCE_PREPARATION-实例资源准备中,RUNNING-实例运行中,RESOURCE_RELEASE-实例资源释放中,DELETED-实例已释放,RESOURCE_PREPARATION_FAIL-实例资源准备失败,FROZEN_RELEASABLE-可恢复冻结,FROZEN_UNRELEASABLE-不可恢复冻结,RECOVERING-恢复中,DELETING-删除中,SCALING-扩容中,SCALE_FAIL-扩容失败
	Status *LakeFormationInstanceStatus `json:"status,omitempty"`

	// 是否在回收站
	InRecycleBin *bool `json:"in_recycle_bin,omitempty"`
}

func (o LakeFormationInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LakeFormationInstance struct{}"
	}

	return strings.Join([]string{"LakeFormationInstance", string(data)}, " ")
}

type LakeFormationInstanceStatus struct {
	value string
}

type LakeFormationInstanceStatusEnum struct {
	RESOURCE_PREPARATION      LakeFormationInstanceStatus
	RUNNING                   LakeFormationInstanceStatus
	RESOURCE_RELEASE          LakeFormationInstanceStatus
	DELETED                   LakeFormationInstanceStatus
	RESOURCE_PREPARATION_FAIL LakeFormationInstanceStatus
	FROZEN_RELEASABLE         LakeFormationInstanceStatus
	FROZEN_UNRELEASABLE       LakeFormationInstanceStatus
	RECOVERING                LakeFormationInstanceStatus
	DELETING                  LakeFormationInstanceStatus
	SCALING                   LakeFormationInstanceStatus
	SCALE_FAIL                LakeFormationInstanceStatus
}

func GetLakeFormationInstanceStatusEnum() LakeFormationInstanceStatusEnum {
	return LakeFormationInstanceStatusEnum{
		RESOURCE_PREPARATION: LakeFormationInstanceStatus{
			value: "RESOURCE_PREPARATION",
		},
		RUNNING: LakeFormationInstanceStatus{
			value: "RUNNING",
		},
		RESOURCE_RELEASE: LakeFormationInstanceStatus{
			value: "RESOURCE_RELEASE",
		},
		DELETED: LakeFormationInstanceStatus{
			value: "DELETED",
		},
		RESOURCE_PREPARATION_FAIL: LakeFormationInstanceStatus{
			value: "RESOURCE_PREPARATION_FAIL",
		},
		FROZEN_RELEASABLE: LakeFormationInstanceStatus{
			value: "FROZEN_RELEASABLE",
		},
		FROZEN_UNRELEASABLE: LakeFormationInstanceStatus{
			value: "FROZEN_UNRELEASABLE",
		},
		RECOVERING: LakeFormationInstanceStatus{
			value: "RECOVERING",
		},
		DELETING: LakeFormationInstanceStatus{
			value: "DELETING",
		},
		SCALING: LakeFormationInstanceStatus{
			value: "SCALING",
		},
		SCALE_FAIL: LakeFormationInstanceStatus{
			value: "SCALE_FAIL",
		},
	}
}

func (c LakeFormationInstanceStatus) Value() string {
	return c.value
}

func (c LakeFormationInstanceStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LakeFormationInstanceStatus) UnmarshalJSON(b []byte) error {
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
