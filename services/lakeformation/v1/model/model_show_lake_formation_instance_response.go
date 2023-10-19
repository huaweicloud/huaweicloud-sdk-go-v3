package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ShowLakeFormationInstanceResponse Response Object
type ShowLakeFormationInstanceResponse struct {

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
	Status *ShowLakeFormationInstanceResponseStatus `json:"status,omitempty"`

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
	ChargeMode *ShowLakeFormationInstanceResponseChargeMode `json:"charge_mode,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowLakeFormationInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLakeFormationInstanceResponse struct{}"
	}

	return strings.Join([]string{"ShowLakeFormationInstanceResponse", string(data)}, " ")
}

type ShowLakeFormationInstanceResponseStatus struct {
	value string
}

type ShowLakeFormationInstanceResponseStatusEnum struct {
	RESOURCE_PREPARATION      ShowLakeFormationInstanceResponseStatus
	RUNNING                   ShowLakeFormationInstanceResponseStatus
	RESOURCE_RELEASE          ShowLakeFormationInstanceResponseStatus
	DELETED                   ShowLakeFormationInstanceResponseStatus
	RESOURCE_PREPARATION_FAIL ShowLakeFormationInstanceResponseStatus
	FROZEN_RELEASABLE         ShowLakeFormationInstanceResponseStatus
	FROZEN_UNRELEASABLE       ShowLakeFormationInstanceResponseStatus
	RECOVERING                ShowLakeFormationInstanceResponseStatus
	DELETING                  ShowLakeFormationInstanceResponseStatus
	SCALING                   ShowLakeFormationInstanceResponseStatus
	SCALE_FAIL                ShowLakeFormationInstanceResponseStatus
}

func GetShowLakeFormationInstanceResponseStatusEnum() ShowLakeFormationInstanceResponseStatusEnum {
	return ShowLakeFormationInstanceResponseStatusEnum{
		RESOURCE_PREPARATION: ShowLakeFormationInstanceResponseStatus{
			value: "RESOURCE_PREPARATION",
		},
		RUNNING: ShowLakeFormationInstanceResponseStatus{
			value: "RUNNING",
		},
		RESOURCE_RELEASE: ShowLakeFormationInstanceResponseStatus{
			value: "RESOURCE_RELEASE",
		},
		DELETED: ShowLakeFormationInstanceResponseStatus{
			value: "DELETED",
		},
		RESOURCE_PREPARATION_FAIL: ShowLakeFormationInstanceResponseStatus{
			value: "RESOURCE_PREPARATION_FAIL",
		},
		FROZEN_RELEASABLE: ShowLakeFormationInstanceResponseStatus{
			value: "FROZEN_RELEASABLE",
		},
		FROZEN_UNRELEASABLE: ShowLakeFormationInstanceResponseStatus{
			value: "FROZEN_UNRELEASABLE",
		},
		RECOVERING: ShowLakeFormationInstanceResponseStatus{
			value: "RECOVERING",
		},
		DELETING: ShowLakeFormationInstanceResponseStatus{
			value: "DELETING",
		},
		SCALING: ShowLakeFormationInstanceResponseStatus{
			value: "SCALING",
		},
		SCALE_FAIL: ShowLakeFormationInstanceResponseStatus{
			value: "SCALE_FAIL",
		},
	}
}

func (c ShowLakeFormationInstanceResponseStatus) Value() string {
	return c.value
}

func (c ShowLakeFormationInstanceResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowLakeFormationInstanceResponseStatus) UnmarshalJSON(b []byte) error {
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

type ShowLakeFormationInstanceResponseChargeMode struct {
	value string
}

type ShowLakeFormationInstanceResponseChargeModeEnum struct {
	PRE_PAID  ShowLakeFormationInstanceResponseChargeMode
	POST_PAID ShowLakeFormationInstanceResponseChargeMode
}

func GetShowLakeFormationInstanceResponseChargeModeEnum() ShowLakeFormationInstanceResponseChargeModeEnum {
	return ShowLakeFormationInstanceResponseChargeModeEnum{
		PRE_PAID: ShowLakeFormationInstanceResponseChargeMode{
			value: "prePaid",
		},
		POST_PAID: ShowLakeFormationInstanceResponseChargeMode{
			value: "postPaid",
		},
	}
}

func (c ShowLakeFormationInstanceResponseChargeMode) Value() string {
	return c.value
}

func (c ShowLakeFormationInstanceResponseChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowLakeFormationInstanceResponseChargeMode) UnmarshalJSON(b []byte) error {
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
