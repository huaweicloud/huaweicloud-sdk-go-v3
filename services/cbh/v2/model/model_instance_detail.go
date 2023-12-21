package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// InstanceDetail CBH实例详情
type InstanceDetail struct {

	// 云堡垒机实例名称。
	Name string `json:"name"`

	// 云堡垒机服务器id。
	ServerId string `json:"server_id"`

	// 云堡垒机实例id。
	InstanceId string `json:"instance_id"`

	// 云堡垒机实例是否可以扩容。 - true：是 - false：否
	AlterPermit bool `json:"alter_permit"`

	// 项目ID。
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 云堡垒机实例订购周期数。
	PeriodNum string `json:"period_num"`

	// 云堡垒机实例开始时间，使用时间戳格式表示。
	StartTime string `json:"start_time"`

	// 云堡垒机实例结束时间，使用时间戳格式表示。
	EndTime string `json:"end_time"`

	// 云堡垒机实例创建时间，使用UTC时间表示。
	CreatedTime string `json:"created_time"`

	// 云堡垒机实例升级定时时间，使用时间戳格式表示。
	UpgradeTime *int64 `json:"upgrade_time,omitempty"`

	// 云堡垒机实例是否可以升级。 - OLD：当前已是最新版本 - NEW：可以升级小版本 - CROSS_OS：可以跨版本升级 - ROLLBACK：可以回滚
	Update InstanceDetailUpdate `json:"update"`

	// 云堡垒机实例当前版本。
	BastionVersion string `json:"bastion_version"`

	AzInfo *InstanceDetailAzInfo `json:"az_info"`

	StatusInfo *InstanceDetailStatusInfo `json:"status_info"`

	ResourceInfo *InstanceDetailResourceInfo `json:"resource_info"`

	Network *InstanceDetailNetwork `json:"network"`

	HaInfo *InstanceDetailHaInfo `json:"ha_info,omitempty"`
}

func (o InstanceDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceDetail struct{}"
	}

	return strings.Join([]string{"InstanceDetail", string(data)}, " ")
}

type InstanceDetailUpdate struct {
	value string
}

type InstanceDetailUpdateEnum struct {
	OLD      InstanceDetailUpdate
	NEW      InstanceDetailUpdate
	CROSS_OS InstanceDetailUpdate
	ROLLBACK InstanceDetailUpdate
}

func GetInstanceDetailUpdateEnum() InstanceDetailUpdateEnum {
	return InstanceDetailUpdateEnum{
		OLD: InstanceDetailUpdate{
			value: "OLD",
		},
		NEW: InstanceDetailUpdate{
			value: "NEW",
		},
		CROSS_OS: InstanceDetailUpdate{
			value: "CROSS_OS",
		},
		ROLLBACK: InstanceDetailUpdate{
			value: "ROLLBACK",
		},
	}
}

func (c InstanceDetailUpdate) Value() string {
	return c.value
}

func (c InstanceDetailUpdate) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceDetailUpdate) UnmarshalJSON(b []byte) error {
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
