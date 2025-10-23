package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ResourceChangeLevelItem struct {

	// 资源ID
	Id string `json:"id"`

	// 结果：succeed-成功，failed-失败
	Result ResourceChangeLevelItemResult `json:"result"`

	// 错误信息
	Message *string `json:"message,omitempty"`
}

func (o ResourceChangeLevelItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceChangeLevelItem struct{}"
	}

	return strings.Join([]string{"ResourceChangeLevelItem", string(data)}, " ")
}

type ResourceChangeLevelItemResult struct {
	value string
}

type ResourceChangeLevelItemResultEnum struct {
	SUCCEED ResourceChangeLevelItemResult
	FAILED  ResourceChangeLevelItemResult
}

func GetResourceChangeLevelItemResultEnum() ResourceChangeLevelItemResultEnum {
	return ResourceChangeLevelItemResultEnum{
		SUCCEED: ResourceChangeLevelItemResult{
			value: "succeed",
		},
		FAILED: ResourceChangeLevelItemResult{
			value: "failed",
		},
	}
}

func (c ResourceChangeLevelItemResult) Value() string {
	return c.value
}

func (c ResourceChangeLevelItemResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceChangeLevelItemResult) UnmarshalJSON(b []byte) error {
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
