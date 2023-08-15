package model

import (
	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type MetaDomain struct {

	// 域名ID。
	Id *string `json:"id,omitempty"`

	// 域名名称。
	Name *string `json:"name,omitempty"`

	// 区域ID。
	ZoneId *string `json:"zone_id,omitempty"`

	// 域名类型，默认是公网域名public
	ZoneType *MetaDomainZoneType `json:"zone_type,omitempty"`

	// 创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`
}

func (o MetaDomain) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetaDomain struct{}"
	}

	return strings.Join([]string{"MetaDomain", string(data)}, " ")
}

type MetaDomainZoneType struct {
	value string
}

type MetaDomainZoneTypeEnum struct {
	PRIVATE MetaDomainZoneType
	PUBLIC  MetaDomainZoneType
}

func GetMetaDomainZoneTypeEnum() MetaDomainZoneTypeEnum {
	return MetaDomainZoneTypeEnum{
		PRIVATE: MetaDomainZoneType{
			value: "private",
		},
		PUBLIC: MetaDomainZoneType{
			value: "public",
		},
	}
}

func (c MetaDomainZoneType) Value() string {
	return c.value
}

func (c MetaDomainZoneType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MetaDomainZoneType) UnmarshalJSON(b []byte) error {
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
