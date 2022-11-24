package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Agency struct {

	// API版本。
	ApiVersion AgencyApiVersion `json:"api_version"`

	// 资源类型。
	Kind AgencyKind `json:"kind"`

	Metadata *AgencyMetadata `json:"metadata,omitempty"`
}

func (o Agency) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Agency struct{}"
	}

	return strings.Join([]string{"Agency", string(data)}, " ")
}

type AgencyApiVersion struct {
	value string
}

type AgencyApiVersionEnum struct {
	V1 AgencyApiVersion
}

func GetAgencyApiVersionEnum() AgencyApiVersionEnum {
	return AgencyApiVersionEnum{
		V1: AgencyApiVersion{
			value: "v1",
		},
	}
}

func (c AgencyApiVersion) Value() string {
	return c.value
}

func (c AgencyApiVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AgencyApiVersion) UnmarshalJSON(b []byte) error {
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

type AgencyKind struct {
	value string
}

type AgencyKindEnum struct {
	AGENCY AgencyKind
}

func GetAgencyKindEnum() AgencyKindEnum {
	return AgencyKindEnum{
		AGENCY: AgencyKind{
			value: "Agency",
		},
	}
}

func (c AgencyKind) Value() string {
	return c.value
}

func (c AgencyKind) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AgencyKind) UnmarshalJSON(b []byte) error {
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
