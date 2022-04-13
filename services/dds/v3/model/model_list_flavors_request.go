package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type ListFlavorsRequest struct {
	// 实例所在区域。

	Region *string `json:"region,omitempty"`
	// 数据库版本类型。取值为“DDS-Community”。

	EngineName *ListFlavorsRequestEngineName `json:"engine_name,omitempty"`
}

func (o ListFlavorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ListFlavorsRequest", string(data)}, " ")
}

type ListFlavorsRequestEngineName struct {
	value string
}

type ListFlavorsRequestEngineNameEnum struct {
	DDS_COMMUNITY ListFlavorsRequestEngineName
	DDS_ENHANCED  ListFlavorsRequestEngineName
}

func GetListFlavorsRequestEngineNameEnum() ListFlavorsRequestEngineNameEnum {
	return ListFlavorsRequestEngineNameEnum{
		DDS_COMMUNITY: ListFlavorsRequestEngineName{
			value: "DDS-Community",
		},
		DDS_ENHANCED: ListFlavorsRequestEngineName{
			value: "DDS-Enhanced",
		},
	}
}

func (c ListFlavorsRequestEngineName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFlavorsRequestEngineName) UnmarshalJSON(b []byte) error {
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
