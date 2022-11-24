package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowAgencyResponse struct {

	// API版本。
	ApiVersion *ShowAgencyResponseApiVersion `json:"api_version,omitempty"`

	// 资源类型。
	Kind *ShowAgencyResponseKind `json:"kind,omitempty"`

	// 已授权委托列表。
	Items          *[]AgencyItem `json:"items,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowAgencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgencyResponse struct{}"
	}

	return strings.Join([]string{"ShowAgencyResponse", string(data)}, " ")
}

type ShowAgencyResponseApiVersion struct {
	value string
}

type ShowAgencyResponseApiVersionEnum struct {
	V1 ShowAgencyResponseApiVersion
}

func GetShowAgencyResponseApiVersionEnum() ShowAgencyResponseApiVersionEnum {
	return ShowAgencyResponseApiVersionEnum{
		V1: ShowAgencyResponseApiVersion{
			value: "v1",
		},
	}
}

func (c ShowAgencyResponseApiVersion) Value() string {
	return c.value
}

func (c ShowAgencyResponseApiVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAgencyResponseApiVersion) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ShowAgencyResponseKind struct {
	value string
}

type ShowAgencyResponseKindEnum struct {
	AGENCY ShowAgencyResponseKind
}

func GetShowAgencyResponseKindEnum() ShowAgencyResponseKindEnum {
	return ShowAgencyResponseKindEnum{
		AGENCY: ShowAgencyResponseKind{
			value: "Agency",
		},
	}
}

func (c ShowAgencyResponseKind) Value() string {
	return c.value
}

func (c ShowAgencyResponseKind) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAgencyResponseKind) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
