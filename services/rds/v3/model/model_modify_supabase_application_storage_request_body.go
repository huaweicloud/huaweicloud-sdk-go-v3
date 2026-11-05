package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ModifySupabaseApplicationStorageRequestBody Storage修改请求。enabled=false时bucket_name/ak/sk必须全空；enabled=true时全必填。
type ModifySupabaseApplicationStorageRequestBody struct {
	Enabled bool `json:"enabled"`

	Type *ModifySupabaseApplicationStorageRequestBodyType `json:"type,omitempty"`

	BucketName *string `json:"bucket_name,omitempty"`

	Ak *string `json:"ak,omitempty"`

	// show不返回，部分更新省略即保留。
	Sk *string `json:"sk,omitempty"`
}

func (o ModifySupabaseApplicationStorageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifySupabaseApplicationStorageRequestBody struct{}"
	}

	return strings.Join([]string{"ModifySupabaseApplicationStorageRequestBody", string(data)}, " ")
}

type ModifySupabaseApplicationStorageRequestBodyType struct {
	value string
}

type ModifySupabaseApplicationStorageRequestBodyTypeEnum struct {
	OBS ModifySupabaseApplicationStorageRequestBodyType
}

func GetModifySupabaseApplicationStorageRequestBodyTypeEnum() ModifySupabaseApplicationStorageRequestBodyTypeEnum {
	return ModifySupabaseApplicationStorageRequestBodyTypeEnum{
		OBS: ModifySupabaseApplicationStorageRequestBodyType{
			value: "obs",
		},
	}
}

func (c ModifySupabaseApplicationStorageRequestBodyType) Value() string {
	return c.value
}

func (c ModifySupabaseApplicationStorageRequestBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifySupabaseApplicationStorageRequestBodyType) UnmarshalJSON(b []byte) error {
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
