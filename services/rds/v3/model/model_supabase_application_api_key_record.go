package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SupabaseApplicationApiKeyRecord ApiKey元数据记录（不含明文）。
type SupabaseApplicationApiKeyRecord struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Remark *string `json:"remark,omitempty"`

	Status *SupabaseApplicationApiKeyRecordStatus `json:"status,omitempty"`

	Type *SupabaseApplicationApiKeyRecordType `json:"type,omitempty"`

	CreatedUser *string `json:"created_user,omitempty"`

	UpdatedUser *string `json:"updated_user,omitempty"`

	CreatedTime *int64 `json:"created_time,omitempty"`

	UpdatedTime *int64 `json:"updated_time,omitempty"`
}

func (o SupabaseApplicationApiKeyRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupabaseApplicationApiKeyRecord struct{}"
	}

	return strings.Join([]string{"SupabaseApplicationApiKeyRecord", string(data)}, " ")
}

type SupabaseApplicationApiKeyRecordStatus struct {
	value string
}

type SupabaseApplicationApiKeyRecordStatusEnum struct {
	NORMAL     SupabaseApplicationApiKeyRecordStatus
	CREATING   SupabaseApplicationApiKeyRecordStatus
	CREATEFAIL SupabaseApplicationApiKeyRecordStatus
	DELETED    SupabaseApplicationApiKeyRecordStatus
}

func GetSupabaseApplicationApiKeyRecordStatusEnum() SupabaseApplicationApiKeyRecordStatusEnum {
	return SupabaseApplicationApiKeyRecordStatusEnum{
		NORMAL: SupabaseApplicationApiKeyRecordStatus{
			value: "normal",
		},
		CREATING: SupabaseApplicationApiKeyRecordStatus{
			value: "creating",
		},
		CREATEFAIL: SupabaseApplicationApiKeyRecordStatus{
			value: "createfail",
		},
		DELETED: SupabaseApplicationApiKeyRecordStatus{
			value: "deleted",
		},
	}
}

func (c SupabaseApplicationApiKeyRecordStatus) Value() string {
	return c.value
}

func (c SupabaseApplicationApiKeyRecordStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SupabaseApplicationApiKeyRecordStatus) UnmarshalJSON(b []byte) error {
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

type SupabaseApplicationApiKeyRecordType struct {
	value string
}

type SupabaseApplicationApiKeyRecordTypeEnum struct {
	PUBLISHABLE  SupabaseApplicationApiKeyRecordType
	SECRET       SupabaseApplicationApiKeyRecordType
	ANON         SupabaseApplicationApiKeyRecordType
	SERVICE_ROLE SupabaseApplicationApiKeyRecordType
}

func GetSupabaseApplicationApiKeyRecordTypeEnum() SupabaseApplicationApiKeyRecordTypeEnum {
	return SupabaseApplicationApiKeyRecordTypeEnum{
		PUBLISHABLE: SupabaseApplicationApiKeyRecordType{
			value: "publishable",
		},
		SECRET: SupabaseApplicationApiKeyRecordType{
			value: "secret",
		},
		ANON: SupabaseApplicationApiKeyRecordType{
			value: "anon",
		},
		SERVICE_ROLE: SupabaseApplicationApiKeyRecordType{
			value: "service_role",
		},
	}
}

func (c SupabaseApplicationApiKeyRecordType) Value() string {
	return c.value
}

func (c SupabaseApplicationApiKeyRecordType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SupabaseApplicationApiKeyRecordType) UnmarshalJSON(b []byte) error {
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
