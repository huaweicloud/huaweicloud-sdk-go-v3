package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateBccPolicyRequestBodyPolicy policy
type CreateBccPolicyRequestBodyPolicy struct {

	// protection_type
	ProtectionType CreateBccPolicyRequestBodyPolicyProtectionType `json:"protection_type"`

	// enabled
	Enabled bool `json:"enabled"`

	// 策略名称，取值长度范围[2,64]。
	Name string `json:"name"`

	// 区域ID
	RegionId string `json:"region_id"`

	// 是否启用加密
	Cyber bool `json:"cyber"`

	// 存储类型，obs,evs,sfs，默认obs
	StorageType CreateBccPolicyRequestBodyPolicyStorageType `json:"storage_type"`

	// 策略类型,cbr_local_policy,cbr_remote_policy,rds_local_policy,rds_remote_policy
	PolicyType CreateBccPolicyRequestBodyPolicyPolicyType `json:"policy_type"`

	// 保护服务类型，如:ecs, evs, rds
	ServiceType string `json:"service_type"`

	CbrPolicyDetail *CreateBccPolicyRequestBodyPolicyCbrPolicyDetail `json:"cbr_policy_detail"`
}

func (o CreateBccPolicyRequestBodyPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBccPolicyRequestBodyPolicy struct{}"
	}

	return strings.Join([]string{"CreateBccPolicyRequestBodyPolicy", string(data)}, " ")
}

type CreateBccPolicyRequestBodyPolicyProtectionType struct {
	value string
}

type CreateBccPolicyRequestBodyPolicyProtectionTypeEnum struct {
	BACKUP            CreateBccPolicyRequestBodyPolicyProtectionType
	DISASTER_RECOVERY CreateBccPolicyRequestBodyPolicyProtectionType
}

func GetCreateBccPolicyRequestBodyPolicyProtectionTypeEnum() CreateBccPolicyRequestBodyPolicyProtectionTypeEnum {
	return CreateBccPolicyRequestBodyPolicyProtectionTypeEnum{
		BACKUP: CreateBccPolicyRequestBodyPolicyProtectionType{
			value: "backup",
		},
		DISASTER_RECOVERY: CreateBccPolicyRequestBodyPolicyProtectionType{
			value: "disaster_recovery",
		},
	}
}

func (c CreateBccPolicyRequestBodyPolicyProtectionType) Value() string {
	return c.value
}

func (c CreateBccPolicyRequestBodyPolicyProtectionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateBccPolicyRequestBodyPolicyProtectionType) UnmarshalJSON(b []byte) error {
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

type CreateBccPolicyRequestBodyPolicyStorageType struct {
	value string
}

type CreateBccPolicyRequestBodyPolicyStorageTypeEnum struct {
	OBS CreateBccPolicyRequestBodyPolicyStorageType
	EVS CreateBccPolicyRequestBodyPolicyStorageType
	SFS CreateBccPolicyRequestBodyPolicyStorageType
}

func GetCreateBccPolicyRequestBodyPolicyStorageTypeEnum() CreateBccPolicyRequestBodyPolicyStorageTypeEnum {
	return CreateBccPolicyRequestBodyPolicyStorageTypeEnum{
		OBS: CreateBccPolicyRequestBodyPolicyStorageType{
			value: "obs",
		},
		EVS: CreateBccPolicyRequestBodyPolicyStorageType{
			value: "evs",
		},
		SFS: CreateBccPolicyRequestBodyPolicyStorageType{
			value: "sfs",
		},
	}
}

func (c CreateBccPolicyRequestBodyPolicyStorageType) Value() string {
	return c.value
}

func (c CreateBccPolicyRequestBodyPolicyStorageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateBccPolicyRequestBodyPolicyStorageType) UnmarshalJSON(b []byte) error {
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

type CreateBccPolicyRequestBodyPolicyPolicyType struct {
	value string
}

type CreateBccPolicyRequestBodyPolicyPolicyTypeEnum struct {
	CBR_LOCAL_POLICY  CreateBccPolicyRequestBodyPolicyPolicyType
	CBR_REMOTE_POLICY CreateBccPolicyRequestBodyPolicyPolicyType
	RDS_LOCAL_POLICY  CreateBccPolicyRequestBodyPolicyPolicyType
	RDS_REMOTE_POLICY CreateBccPolicyRequestBodyPolicyPolicyType
}

func GetCreateBccPolicyRequestBodyPolicyPolicyTypeEnum() CreateBccPolicyRequestBodyPolicyPolicyTypeEnum {
	return CreateBccPolicyRequestBodyPolicyPolicyTypeEnum{
		CBR_LOCAL_POLICY: CreateBccPolicyRequestBodyPolicyPolicyType{
			value: "cbr_local_policy",
		},
		CBR_REMOTE_POLICY: CreateBccPolicyRequestBodyPolicyPolicyType{
			value: "cbr_remote_policy",
		},
		RDS_LOCAL_POLICY: CreateBccPolicyRequestBodyPolicyPolicyType{
			value: "rds_local_policy",
		},
		RDS_REMOTE_POLICY: CreateBccPolicyRequestBodyPolicyPolicyType{
			value: "rds_remote_policy",
		},
	}
}

func (c CreateBccPolicyRequestBodyPolicyPolicyType) Value() string {
	return c.value
}

func (c CreateBccPolicyRequestBodyPolicyPolicyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateBccPolicyRequestBodyPolicyPolicyType) UnmarshalJSON(b []byte) error {
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
