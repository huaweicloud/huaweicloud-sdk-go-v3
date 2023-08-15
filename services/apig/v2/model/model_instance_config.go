package model

import (
	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type InstanceConfig struct {

	// 配额编号
	ConfigId *string `json:"config_id,omitempty"`

	// 配额名称
	ConfigName *InstanceConfigConfigName `json:"config_name,omitempty"`

	// 配额值  当前实例所在租户该配额对应的数量
	ConfigValue *string `json:"config_value,omitempty"`

	// 配额创建时间
	ConfigTime *sdktime.SdkTime `json:"config_time,omitempty"`

	// 配额描述 - INSTANCE_NUM_LIMIT：租户可以创建的实例个数限制
	Remark *string `json:"remark,omitempty"`
}

func (o InstanceConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceConfig struct{}"
	}

	return strings.Join([]string{"InstanceConfig", string(data)}, " ")
}

type InstanceConfigConfigName struct {
	value string
}

type InstanceConfigConfigNameEnum struct {
	INSTANCE_NUM_LIMIT InstanceConfigConfigName
}

func GetInstanceConfigConfigNameEnum() InstanceConfigConfigNameEnum {
	return InstanceConfigConfigNameEnum{
		INSTANCE_NUM_LIMIT: InstanceConfigConfigName{
			value: "INSTANCE_NUM_LIMIT",
		},
	}
}

func (c InstanceConfigConfigName) Value() string {
	return c.value
}

func (c InstanceConfigConfigName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceConfigConfigName) UnmarshalJSON(b []byte) error {
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
