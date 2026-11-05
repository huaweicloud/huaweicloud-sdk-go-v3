package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SupabaseStorageComponentConfig Supabase Storage存储组件对接配置。
type SupabaseStorageComponentConfig struct {

	// 是否启用Storage组件。enabled=false时bucket_name/ak/sk必须全空；enabled=true时全必填。
	Enabled bool `json:"enabled"`

	// 对象存储类型。
	Type *SupabaseStorageComponentConfigType `json:"type,omitempty"`

	// 桶名称，3~64字符。enabled=true时必填。
	BucketName *string `json:"bucket_name,omitempty"`

	// 访问密钥ID，1~32字符。enabled=true时必填。
	Ak *string `json:"ak,omitempty"`

	// 秘密访问密钥，1~64字符。enabled=true时必填。
	Sk *string `json:"sk,omitempty"`
}

func (o SupabaseStorageComponentConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupabaseStorageComponentConfig struct{}"
	}

	return strings.Join([]string{"SupabaseStorageComponentConfig", string(data)}, " ")
}

type SupabaseStorageComponentConfigType struct {
	value string
}

type SupabaseStorageComponentConfigTypeEnum struct {
	OBS SupabaseStorageComponentConfigType
}

func GetSupabaseStorageComponentConfigTypeEnum() SupabaseStorageComponentConfigTypeEnum {
	return SupabaseStorageComponentConfigTypeEnum{
		OBS: SupabaseStorageComponentConfigType{
			value: "obs",
		},
	}
}

func (c SupabaseStorageComponentConfigType) Value() string {
	return c.value
}

func (c SupabaseStorageComponentConfigType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SupabaseStorageComponentConfigType) UnmarshalJSON(b []byte) error {
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
