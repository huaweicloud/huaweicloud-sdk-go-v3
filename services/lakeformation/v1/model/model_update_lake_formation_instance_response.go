package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// UpdateLakeFormationInstanceResponse Response Object
type UpdateLakeFormationInstanceResponse struct {

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
	Status *UpdateLakeFormationInstanceResponseStatus `json:"status,omitempty"`

	// 是否在回收站
	InRecycleBin   *bool `json:"in_recycle_bin,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o UpdateLakeFormationInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLakeFormationInstanceResponse struct{}"
	}

	return strings.Join([]string{"UpdateLakeFormationInstanceResponse", string(data)}, " ")
}

type UpdateLakeFormationInstanceResponseStatus struct {
	value string
}

type UpdateLakeFormationInstanceResponseStatusEnum struct {
	RESOURCE_PREPARATION      UpdateLakeFormationInstanceResponseStatus
	RUNNING                   UpdateLakeFormationInstanceResponseStatus
	RESOURCE_RELEASE          UpdateLakeFormationInstanceResponseStatus
	DELETED                   UpdateLakeFormationInstanceResponseStatus
	RESOURCE_PREPARATION_FAIL UpdateLakeFormationInstanceResponseStatus
	FROZEN_RELEASABLE         UpdateLakeFormationInstanceResponseStatus
	FROZEN_UNRELEASABLE       UpdateLakeFormationInstanceResponseStatus
	RECOVERING                UpdateLakeFormationInstanceResponseStatus
	DELETING                  UpdateLakeFormationInstanceResponseStatus
	SCALING                   UpdateLakeFormationInstanceResponseStatus
	SCALE_FAIL                UpdateLakeFormationInstanceResponseStatus
}

func GetUpdateLakeFormationInstanceResponseStatusEnum() UpdateLakeFormationInstanceResponseStatusEnum {
	return UpdateLakeFormationInstanceResponseStatusEnum{
		RESOURCE_PREPARATION: UpdateLakeFormationInstanceResponseStatus{
			value: "RESOURCE_PREPARATION",
		},
		RUNNING: UpdateLakeFormationInstanceResponseStatus{
			value: "RUNNING",
		},
		RESOURCE_RELEASE: UpdateLakeFormationInstanceResponseStatus{
			value: "RESOURCE_RELEASE",
		},
		DELETED: UpdateLakeFormationInstanceResponseStatus{
			value: "DELETED",
		},
		RESOURCE_PREPARATION_FAIL: UpdateLakeFormationInstanceResponseStatus{
			value: "RESOURCE_PREPARATION_FAIL",
		},
		FROZEN_RELEASABLE: UpdateLakeFormationInstanceResponseStatus{
			value: "FROZEN_RELEASABLE",
		},
		FROZEN_UNRELEASABLE: UpdateLakeFormationInstanceResponseStatus{
			value: "FROZEN_UNRELEASABLE",
		},
		RECOVERING: UpdateLakeFormationInstanceResponseStatus{
			value: "RECOVERING",
		},
		DELETING: UpdateLakeFormationInstanceResponseStatus{
			value: "DELETING",
		},
		SCALING: UpdateLakeFormationInstanceResponseStatus{
			value: "SCALING",
		},
		SCALE_FAIL: UpdateLakeFormationInstanceResponseStatus{
			value: "SCALE_FAIL",
		},
	}
}

func (c UpdateLakeFormationInstanceResponseStatus) Value() string {
	return c.value
}

func (c UpdateLakeFormationInstanceResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateLakeFormationInstanceResponseStatus) UnmarshalJSON(b []byte) error {
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
