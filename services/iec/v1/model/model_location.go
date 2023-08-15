package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Location 位置信息
type Location struct {

	// 站点ID。
	SiteId *string `json:"site_id,omitempty"`

	// 所在大区。
	Area *string `json:"area,omitempty"`

	// 所属省份英文名称。 大小写通用，皆支持
	Province *string `json:"province,omitempty"`

	// 所在城市英文名称。
	City *string `json:"city,omitempty"`

	// 所属运营商。
	Operator *LocationOperator `json:"operator,omitempty"`

	// 线路ID。多线路场景下，创建的弹性公网IP在该线路下。
	PoolId *string `json:"pool_id,omitempty"`

	// 站点需要发放的资源(组)总数。
	StackCount *int32 `json:"stack_count,omitempty"`

	// 城市简称。
	CityShortName *string `json:"city_short_name,omitempty"`
}

func (o Location) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Location struct{}"
	}

	return strings.Join([]string{"Location", string(data)}, " ")
}

type LocationOperator struct {
	value string
}

type LocationOperatorEnum struct {
	CHINAMOBILE  LocationOperator
	CHINAUNICOM  LocationOperator
	CHINATELECOM LocationOperator
}

func GetLocationOperatorEnum() LocationOperatorEnum {
	return LocationOperatorEnum{
		CHINAMOBILE: LocationOperator{
			value: "chinamobile",
		},
		CHINAUNICOM: LocationOperator{
			value: "chinaunicom",
		},
		CHINATELECOM: LocationOperator{
			value: "chinatelecom",
		},
	}
}

func (c LocationOperator) Value() string {
	return c.value
}

func (c LocationOperator) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LocationOperator) UnmarshalJSON(b []byte) error {
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
