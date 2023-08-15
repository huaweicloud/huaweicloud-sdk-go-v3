package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Billing struct {

	// 已分配容量，单位GB
	Allocated int32 `json:"allocated"`

	// 创建模式
	ChargingMode BillingChargingMode `json:"charging_mode"`

	// 云平台
	CloudType *BillingCloudType `json:"cloud_type,omitempty"`

	// 崩溃一致性（crash_consistent）或应用一致性（app_consistent）
	ConsistentLevel BillingConsistentLevel `json:"consistent_level"`

	// 对象类型：云服务器（server），云硬盘（disk），文件系统（turbo），云桌面（workspace），VMware（vmware），关系型数据库（rds），文件（file）。
	ObjectType *BillingObjectType `json:"object_type,omitempty"`

	// 订单ID
	OrderId *string `json:"order_id,omitempty"`

	// 产品ID
	ProductId *string `json:"product_id,omitempty"`

	// 保护类型
	ProtectType BillingProtectType `json:"protect_type"`

	// 容量，单位GB
	Size int32 `json:"size"`

	// 规格编码: 云服务备份存储库:vault.backup.server.normal;云硬盘备份存储库:vault.backup.volume.normal;文件备份存储库:vault.backup.turbo.normal;数据库备份存储库:vault.backup.database.normal;混合云备份存储库:vault.hybrid.server.normal;复制备份存储库:vault.replication.server.normal
	SpecCode BillingSpecCode `json:"spec_code"`

	// 存储库状态
	Status BillingStatus `json:"status"`

	// 存储库桶名
	StorageUnit *string `json:"storage_unit,omitempty"`

	// 已使用容量，单位MB
	Used int32 `json:"used"`

	// 冻结场景
	FrozenScene *string `json:"frozen_scene,omitempty"`

	// 存储库多az属性
	IsMultiAz *bool `json:"is_multi_az,omitempty"`
}

func (o Billing) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Billing struct{}"
	}

	return strings.Join([]string{"Billing", string(data)}, " ")
}

type BillingChargingMode struct {
	value string
}

type BillingChargingModeEnum struct {
	PRE_PAID  BillingChargingMode
	POST_PAID BillingChargingMode
}

func GetBillingChargingModeEnum() BillingChargingModeEnum {
	return BillingChargingModeEnum{
		PRE_PAID: BillingChargingMode{
			value: "pre_paid",
		},
		POST_PAID: BillingChargingMode{
			value: "post_paid",
		},
	}
}

func (c BillingChargingMode) Value() string {
	return c.value
}

func (c BillingChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BillingChargingMode) UnmarshalJSON(b []byte) error {
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

type BillingCloudType struct {
	value string
}

type BillingCloudTypeEnum struct {
	PUBLIC BillingCloudType
	HYBRID BillingCloudType
}

func GetBillingCloudTypeEnum() BillingCloudTypeEnum {
	return BillingCloudTypeEnum{
		PUBLIC: BillingCloudType{
			value: "public",
		},
		HYBRID: BillingCloudType{
			value: "hybrid",
		},
	}
}

func (c BillingCloudType) Value() string {
	return c.value
}

func (c BillingCloudType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BillingCloudType) UnmarshalJSON(b []byte) error {
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

type BillingConsistentLevel struct {
	value string
}

type BillingConsistentLevelEnum struct {
	APP_CONSISTENT   BillingConsistentLevel
	CRASH_CONSISTENT BillingConsistentLevel
}

func GetBillingConsistentLevelEnum() BillingConsistentLevelEnum {
	return BillingConsistentLevelEnum{
		APP_CONSISTENT: BillingConsistentLevel{
			value: "app_consistent",
		},
		CRASH_CONSISTENT: BillingConsistentLevel{
			value: "crash_consistent",
		},
	}
}

func (c BillingConsistentLevel) Value() string {
	return c.value
}

func (c BillingConsistentLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BillingConsistentLevel) UnmarshalJSON(b []byte) error {
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

type BillingObjectType struct {
	value string
}

type BillingObjectTypeEnum struct {
	SERVER    BillingObjectType
	DISK      BillingObjectType
	WORKSPACE BillingObjectType
	VMWARE    BillingObjectType
	RDS       BillingObjectType
	FILE      BillingObjectType
}

func GetBillingObjectTypeEnum() BillingObjectTypeEnum {
	return BillingObjectTypeEnum{
		SERVER: BillingObjectType{
			value: "server",
		},
		DISK: BillingObjectType{
			value: "disk",
		},
		WORKSPACE: BillingObjectType{
			value: "workspace",
		},
		VMWARE: BillingObjectType{
			value: "vmware",
		},
		RDS: BillingObjectType{
			value: "rds",
		},
		FILE: BillingObjectType{
			value: "file",
		},
	}
}

func (c BillingObjectType) Value() string {
	return c.value
}

func (c BillingObjectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BillingObjectType) UnmarshalJSON(b []byte) error {
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

type BillingProtectType struct {
	value string
}

type BillingProtectTypeEnum struct {
	BACKUP      BillingProtectType
	REPLICATION BillingProtectType
	HYBRID      BillingProtectType
}

func GetBillingProtectTypeEnum() BillingProtectTypeEnum {
	return BillingProtectTypeEnum{
		BACKUP: BillingProtectType{
			value: "backup",
		},
		REPLICATION: BillingProtectType{
			value: "replication",
		},
		HYBRID: BillingProtectType{
			value: "hybrid",
		},
	}
}

func (c BillingProtectType) Value() string {
	return c.value
}

func (c BillingProtectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BillingProtectType) UnmarshalJSON(b []byte) error {
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

type BillingSpecCode struct {
	value string
}

type BillingSpecCodeEnum struct {
	VAULT_BACKUP_SERVER_NORMAL BillingSpecCode
	VAULT_BACKUP_VOLUME_NORMAL BillingSpecCode
}

func GetBillingSpecCodeEnum() BillingSpecCodeEnum {
	return BillingSpecCodeEnum{
		VAULT_BACKUP_SERVER_NORMAL: BillingSpecCode{
			value: "vault.backup.server.normal",
		},
		VAULT_BACKUP_VOLUME_NORMAL: BillingSpecCode{
			value: "vault.backup.volume.normal",
		},
	}
}

func (c BillingSpecCode) Value() string {
	return c.value
}

func (c BillingSpecCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BillingSpecCode) UnmarshalJSON(b []byte) error {
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

type BillingStatus struct {
	value string
}

type BillingStatusEnum struct {
	AVAILABLE BillingStatus
	LOCK      BillingStatus
	FROZEN    BillingStatus
	DELETING  BillingStatus
	ERROR     BillingStatus
}

func GetBillingStatusEnum() BillingStatusEnum {
	return BillingStatusEnum{
		AVAILABLE: BillingStatus{
			value: "available",
		},
		LOCK: BillingStatus{
			value: "lock",
		},
		FROZEN: BillingStatus{
			value: "frozen",
		},
		DELETING: BillingStatus{
			value: "deleting",
		},
		ERROR: BillingStatus{
			value: "error",
		},
	}
}

func (c BillingStatus) Value() string {
	return c.value
}

func (c BillingStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BillingStatus) UnmarshalJSON(b []byte) error {
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
