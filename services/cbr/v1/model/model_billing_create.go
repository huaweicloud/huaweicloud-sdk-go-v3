package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 创建参数
type BillingCreate struct {
	// 云平台，公有云或者混合云

	CloudType *BillingCreateCloudType `json:"cloud_type,omitempty"`
	// 规格，崩溃一致性（crash_consistent）或应用一致性（app_consistent）

	ConsistentLevel BillingCreateConsistentLevel `json:"consistent_level"`
	// 对象类型：云服务器（server），云硬盘（disk），文件系统（turbo）。

	ObjectType BillingCreateObjectType `json:"object_type"`
	// 保护类型：备份（backup）、复制(replication)

	ProtectType BillingCreateProtectType `json:"protect_type"`
	// 容量，单位GB

	Size int32 `json:"size"`
	// 创建模式，按需：post_paid，包周期：pre_paid，默认为post_paid

	ChargingMode *BillingCreateChargingMode `json:"charging_mode,omitempty"`
	// 创建类型，按年(year)或者按月(month)

	PeriodType *BillingCreatePeriodType `json:"period_type,omitempty"`
	// 创建类型的数量

	PeriodNum *int32 `json:"period_num,omitempty"`
	// 到期后是否自动续期，默认不续期

	IsAutoRenew *bool `json:"is_auto_renew,omitempty"`
	// 是否自动付费，默认为不自动付费

	IsAutoPay *bool `json:"is_auto_pay,omitempty"`
	// 跳转URL

	ConsoleUrl *string `json:"console_url,omitempty"`

	ExtraInfo *BillbingCreateExtraInfo `json:"extra_info,omitempty"`
}

func (o BillingCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BillingCreate struct{}"
	}

	return strings.Join([]string{"BillingCreate", string(data)}, " ")
}

type BillingCreateCloudType struct {
	value string
}

type BillingCreateCloudTypeEnum struct {
	PUBLIC BillingCreateCloudType
	HYBRID BillingCreateCloudType
}

func GetBillingCreateCloudTypeEnum() BillingCreateCloudTypeEnum {
	return BillingCreateCloudTypeEnum{
		PUBLIC: BillingCreateCloudType{
			value: "public",
		},
		HYBRID: BillingCreateCloudType{
			value: "hybrid",
		},
	}
}

func (c BillingCreateCloudType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BillingCreateCloudType) UnmarshalJSON(b []byte) error {
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

type BillingCreateConsistentLevel struct {
	value string
}

type BillingCreateConsistentLevelEnum struct {
	APP_CONSISTENT   BillingCreateConsistentLevel
	CRASH_CONSISTENT BillingCreateConsistentLevel
}

func GetBillingCreateConsistentLevelEnum() BillingCreateConsistentLevelEnum {
	return BillingCreateConsistentLevelEnum{
		APP_CONSISTENT: BillingCreateConsistentLevel{
			value: "app_consistent",
		},
		CRASH_CONSISTENT: BillingCreateConsistentLevel{
			value: "crash_consistent",
		},
	}
}

func (c BillingCreateConsistentLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BillingCreateConsistentLevel) UnmarshalJSON(b []byte) error {
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

type BillingCreateObjectType struct {
	value string
}

type BillingCreateObjectTypeEnum struct {
	SERVER BillingCreateObjectType
	DISK   BillingCreateObjectType
	TURBO  BillingCreateObjectType
}

func GetBillingCreateObjectTypeEnum() BillingCreateObjectTypeEnum {
	return BillingCreateObjectTypeEnum{
		SERVER: BillingCreateObjectType{
			value: "server",
		},
		DISK: BillingCreateObjectType{
			value: "disk",
		},
		TURBO: BillingCreateObjectType{
			value: "turbo",
		},
	}
}

func (c BillingCreateObjectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BillingCreateObjectType) UnmarshalJSON(b []byte) error {
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

type BillingCreateProtectType struct {
	value string
}

type BillingCreateProtectTypeEnum struct {
	BACKUP      BillingCreateProtectType
	REPLICATION BillingCreateProtectType
}

func GetBillingCreateProtectTypeEnum() BillingCreateProtectTypeEnum {
	return BillingCreateProtectTypeEnum{
		BACKUP: BillingCreateProtectType{
			value: "backup",
		},
		REPLICATION: BillingCreateProtectType{
			value: "replication",
		},
	}
}

func (c BillingCreateProtectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BillingCreateProtectType) UnmarshalJSON(b []byte) error {
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

type BillingCreateChargingMode struct {
	value string
}

type BillingCreateChargingModeEnum struct {
	POST_PAID BillingCreateChargingMode
	PRE_PAID  BillingCreateChargingMode
}

func GetBillingCreateChargingModeEnum() BillingCreateChargingModeEnum {
	return BillingCreateChargingModeEnum{
		POST_PAID: BillingCreateChargingMode{
			value: "post_paid",
		},
		PRE_PAID: BillingCreateChargingMode{
			value: "pre_paid",
		},
	}
}

func (c BillingCreateChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BillingCreateChargingMode) UnmarshalJSON(b []byte) error {
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

type BillingCreatePeriodType struct {
	value string
}

type BillingCreatePeriodTypeEnum struct {
	YEAR  BillingCreatePeriodType
	MONTH BillingCreatePeriodType
}

func GetBillingCreatePeriodTypeEnum() BillingCreatePeriodTypeEnum {
	return BillingCreatePeriodTypeEnum{
		YEAR: BillingCreatePeriodType{
			value: "year",
		},
		MONTH: BillingCreatePeriodType{
			value: "month",
		},
	}
}

func (c BillingCreatePeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BillingCreatePeriodType) UnmarshalJSON(b []byte) error {
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
