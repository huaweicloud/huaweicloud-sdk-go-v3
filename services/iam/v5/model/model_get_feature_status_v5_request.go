package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// GetFeatureStatusV5Request Request Object
type GetFeatureStatusV5Request struct {

	// 功能特征的唯一名称。
	FeatureName GetFeatureStatusV5RequestFeatureName `json:"feature_name"`
}

func (o GetFeatureStatusV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetFeatureStatusV5Request struct{}"
	}

	return strings.Join([]string{"GetFeatureStatusV5Request", string(data)}, " ")
}

type GetFeatureStatusV5RequestFeatureName struct {
	value string
}

type GetFeatureStatusV5RequestFeatureNameEnum struct {
	V5_CONSOLE      GetFeatureStatusV5RequestFeatureName
	ACCESS_ANALYZER GetFeatureStatusV5RequestFeatureName
}

func GetGetFeatureStatusV5RequestFeatureNameEnum() GetFeatureStatusV5RequestFeatureNameEnum {
	return GetFeatureStatusV5RequestFeatureNameEnum{
		V5_CONSOLE: GetFeatureStatusV5RequestFeatureName{
			value: "v5_console",
		},
		ACCESS_ANALYZER: GetFeatureStatusV5RequestFeatureName{
			value: "access_analyzer",
		},
	}
}

func (c GetFeatureStatusV5RequestFeatureName) Value() string {
	return c.value
}

func (c GetFeatureStatusV5RequestFeatureName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetFeatureStatusV5RequestFeatureName) UnmarshalJSON(b []byte) error {
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
