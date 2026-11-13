package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateSupabaseApplicationRequestBody 创建Supabase应用请求体。
type CreateSupabaseApplicationRequestBody struct {

	// Supabase应用名称，4~64个字符。
	Name string `json:"name"`

	BasicConfig *SupabaseAppBasicConfig `json:"basic_config"`

	// Supabase版本号，格式 x.y.z。
	Version string `json:"version"`

	// 部署模式。当前仅支持Single。
	Mode CreateSupabaseApplicationRequestBodyMode `json:"mode"`

	// 规格码，格式 rds.supabase.{...}。
	FlavorRef string `json:"flavor_ref"`

	// 可用区ID。Single模式下为1个可用区。
	AvailabilityZone string `json:"availability_zone"`

	// 指定实例内网IP（可选）。IPv4格式。
	DataVip *string `json:"data_vip,omitempty"`

	// 指定实例所属安全组（可选）。UUID格式。
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	// 默认用户supabase的密码。
	Password string `json:"password"`

	StorageComponent *SupabaseStorageComponentConfig `json:"storage_component"`

	// 实例纳管方式。默认newly_instance。
	ManageType *CreateSupabaseApplicationRequestBodyManageType `json:"manage_type,omitempty"`

	ManagedInstance *SupabaseInstanceCreateRequest `json:"managed_instance,omitempty"`

	ExistingInstance *SupabaseExistingInstanceRequest `json:"existing_instance,omitempty"`
}

func (o CreateSupabaseApplicationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSupabaseApplicationRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSupabaseApplicationRequestBody", string(data)}, " ")
}

type CreateSupabaseApplicationRequestBodyMode struct {
	value string
}

type CreateSupabaseApplicationRequestBodyModeEnum struct {
	SINGLE CreateSupabaseApplicationRequestBodyMode
}

func GetCreateSupabaseApplicationRequestBodyModeEnum() CreateSupabaseApplicationRequestBodyModeEnum {
	return CreateSupabaseApplicationRequestBodyModeEnum{
		SINGLE: CreateSupabaseApplicationRequestBodyMode{
			value: "Single",
		},
	}
}

func (c CreateSupabaseApplicationRequestBodyMode) Value() string {
	return c.value
}

func (c CreateSupabaseApplicationRequestBodyMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSupabaseApplicationRequestBodyMode) UnmarshalJSON(b []byte) error {
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

type CreateSupabaseApplicationRequestBodyManageType struct {
	value string
}

type CreateSupabaseApplicationRequestBodyManageTypeEnum struct {
	NEWLY_INSTANCE   CreateSupabaseApplicationRequestBodyManageType
	EXISTED_INSTANCE CreateSupabaseApplicationRequestBodyManageType
}

func GetCreateSupabaseApplicationRequestBodyManageTypeEnum() CreateSupabaseApplicationRequestBodyManageTypeEnum {
	return CreateSupabaseApplicationRequestBodyManageTypeEnum{
		NEWLY_INSTANCE: CreateSupabaseApplicationRequestBodyManageType{
			value: "newly_instance",
		},
		EXISTED_INSTANCE: CreateSupabaseApplicationRequestBodyManageType{
			value: "existed_instance",
		},
	}
}

func (c CreateSupabaseApplicationRequestBodyManageType) Value() string {
	return c.value
}

func (c CreateSupabaseApplicationRequestBodyManageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSupabaseApplicationRequestBodyManageType) UnmarshalJSON(b []byte) error {
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
