package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 请求数据。
type AgencyMetadata struct {

	// 委托名称。
	Name AgencyMetadataName `json:"name"`
}

func (o AgencyMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyMetadata struct{}"
	}

	return strings.Join([]string{"AgencyMetadata", string(data)}, " ")
}

type AgencyMetadataName struct {
	value string
}

type AgencyMetadataNameEnum struct {
	CAE_TRUST AgencyMetadataName
}

func GetAgencyMetadataNameEnum() AgencyMetadataNameEnum {
	return AgencyMetadataNameEnum{
		CAE_TRUST: AgencyMetadataName{
			value: "cae_trust",
		},
	}
}

func (c AgencyMetadataName) Value() string {
	return c.value
}

func (c AgencyMetadataName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AgencyMetadataName) UnmarshalJSON(b []byte) error {
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
