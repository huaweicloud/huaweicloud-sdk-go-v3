package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateNatMappingConfig NAT映射配置。
type UpdateNatMappingConfig struct {

	// NAT映射类型。 - PORT:端口映射. - HOST:地址映射,最多支持配置10个地址.
	NatMapType *UpdateNatMappingConfigNatMapType `json:"nat_map_type,omitempty"`

	// nat_map_type为PORT时表示端口,取值9443/9445. nat_map_type为HOST时表示接入地址.
	NatMapValue *string `json:"nat_map_value,omitempty"`

	// nat Ip。
	NatIp *string `json:"nat_ip,omitempty"`

	// nat端口,取值范围0-65535。
	NatPort *string `json:"nat_port,omitempty"`

	// vag Ip。
	VagIp *string `json:"vag_ip,omitempty"`

	// 0标识不开启，1表示开启。
	AccessFilterType *int32 `json:"access_filter_type,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 标签对象
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o UpdateNatMappingConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNatMappingConfig struct{}"
	}

	return strings.Join([]string{"UpdateNatMappingConfig", string(data)}, " ")
}

type UpdateNatMappingConfigNatMapType struct {
	value string
}

type UpdateNatMappingConfigNatMapTypeEnum struct {
	PORT UpdateNatMappingConfigNatMapType
	HOST UpdateNatMappingConfigNatMapType
}

func GetUpdateNatMappingConfigNatMapTypeEnum() UpdateNatMappingConfigNatMapTypeEnum {
	return UpdateNatMappingConfigNatMapTypeEnum{
		PORT: UpdateNatMappingConfigNatMapType{
			value: "PORT",
		},
		HOST: UpdateNatMappingConfigNatMapType{
			value: "HOST",
		},
	}
}

func (c UpdateNatMappingConfigNatMapType) Value() string {
	return c.value
}

func (c UpdateNatMappingConfigNatMapType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateNatMappingConfigNatMapType) UnmarshalJSON(b []byte) error {
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
