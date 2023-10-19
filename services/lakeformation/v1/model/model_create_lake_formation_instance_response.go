package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// CreateLakeFormationInstanceResponse Response Object
type CreateLakeFormationInstanceResponse struct {

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
	Status *CreateLakeFormationInstanceResponseStatus `json:"status,omitempty"`

	// 资源准备进度百分比，只有当实例处于资源准备中状态时才会计算并返回该值
	ResourceProgress *int32 `json:"resource_progress,omitempty"`

	// 资源准备预计时长，单位分钟
	ResourceExpectedDuration *int32 `json:"resource_expected_duration,omitempty"`

	// 规格变更进度百分比，只有当实例处于规格变更中状态时才会计算并返回该值
	ScaleProgress *int32 `json:"scale_progress,omitempty"`

	// 规格变更预计时长，单位分钟
	ScaleExpectedDuration *int32 `json:"scale_expected_duration,omitempty"`

	// 是否在回收站
	InRecycleBin *bool `json:"in_recycle_bin,omitempty"`

	// 标签列表
	Tags *[]ResourceTag `json:"tags,omitempty"`

	// 规格信息
	Specs *[]CreateSpec `json:"specs,omitempty"`

	// 计费模式,postPaid=按需计费,prePaid=包周期计费
	ChargeMode *CreateLakeFormationInstanceResponseChargeMode `json:"charge_mode,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateLakeFormationInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLakeFormationInstanceResponse struct{}"
	}

	return strings.Join([]string{"CreateLakeFormationInstanceResponse", string(data)}, " ")
}

type CreateLakeFormationInstanceResponseStatus struct {
	value string
}

type CreateLakeFormationInstanceResponseStatusEnum struct {
	RESOURCE_PREPARATION      CreateLakeFormationInstanceResponseStatus
	RUNNING                   CreateLakeFormationInstanceResponseStatus
	RESOURCE_RELEASE          CreateLakeFormationInstanceResponseStatus
	DELETED                   CreateLakeFormationInstanceResponseStatus
	RESOURCE_PREPARATION_FAIL CreateLakeFormationInstanceResponseStatus
	FROZEN_RELEASABLE         CreateLakeFormationInstanceResponseStatus
	FROZEN_UNRELEASABLE       CreateLakeFormationInstanceResponseStatus
	RECOVERING                CreateLakeFormationInstanceResponseStatus
	DELETING                  CreateLakeFormationInstanceResponseStatus
	SCALING                   CreateLakeFormationInstanceResponseStatus
	SCALE_FAIL                CreateLakeFormationInstanceResponseStatus
}

func GetCreateLakeFormationInstanceResponseStatusEnum() CreateLakeFormationInstanceResponseStatusEnum {
	return CreateLakeFormationInstanceResponseStatusEnum{
		RESOURCE_PREPARATION: CreateLakeFormationInstanceResponseStatus{
			value: "RESOURCE_PREPARATION",
		},
		RUNNING: CreateLakeFormationInstanceResponseStatus{
			value: "RUNNING",
		},
		RESOURCE_RELEASE: CreateLakeFormationInstanceResponseStatus{
			value: "RESOURCE_RELEASE",
		},
		DELETED: CreateLakeFormationInstanceResponseStatus{
			value: "DELETED",
		},
		RESOURCE_PREPARATION_FAIL: CreateLakeFormationInstanceResponseStatus{
			value: "RESOURCE_PREPARATION_FAIL",
		},
		FROZEN_RELEASABLE: CreateLakeFormationInstanceResponseStatus{
			value: "FROZEN_RELEASABLE",
		},
		FROZEN_UNRELEASABLE: CreateLakeFormationInstanceResponseStatus{
			value: "FROZEN_UNRELEASABLE",
		},
		RECOVERING: CreateLakeFormationInstanceResponseStatus{
			value: "RECOVERING",
		},
		DELETING: CreateLakeFormationInstanceResponseStatus{
			value: "DELETING",
		},
		SCALING: CreateLakeFormationInstanceResponseStatus{
			value: "SCALING",
		},
		SCALE_FAIL: CreateLakeFormationInstanceResponseStatus{
			value: "SCALE_FAIL",
		},
	}
}

func (c CreateLakeFormationInstanceResponseStatus) Value() string {
	return c.value
}

func (c CreateLakeFormationInstanceResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLakeFormationInstanceResponseStatus) UnmarshalJSON(b []byte) error {
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

type CreateLakeFormationInstanceResponseChargeMode struct {
	value string
}

type CreateLakeFormationInstanceResponseChargeModeEnum struct {
	PRE_PAID  CreateLakeFormationInstanceResponseChargeMode
	POST_PAID CreateLakeFormationInstanceResponseChargeMode
}

func GetCreateLakeFormationInstanceResponseChargeModeEnum() CreateLakeFormationInstanceResponseChargeModeEnum {
	return CreateLakeFormationInstanceResponseChargeModeEnum{
		PRE_PAID: CreateLakeFormationInstanceResponseChargeMode{
			value: "prePaid",
		},
		POST_PAID: CreateLakeFormationInstanceResponseChargeMode{
			value: "postPaid",
		},
	}
}

func (c CreateLakeFormationInstanceResponseChargeMode) Value() string {
	return c.value
}

func (c CreateLakeFormationInstanceResponseChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLakeFormationInstanceResponseChargeMode) UnmarshalJSON(b []byte) error {
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
