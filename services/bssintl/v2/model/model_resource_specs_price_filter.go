package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ResourceSpecsPriceFilter struct {

	// 过滤条件的key值，必填，支持RESOURCE_SPEC：资源规格编码、CHARGING_MODE：计费模式
	Key ResourceSpecsPriceFilterKey `json:"key"`

	// 过滤条件的value值，必填，不支持模糊查询。当key=CHARGING_MODE时，此处取值如下：PERIOD：包年/包月、ON_DEMAND：按需、ONE_TIME：一次性、ON_DEMAND_PKG：按需套餐包
	Value string `json:"value"`
}

func (o ResourceSpecsPriceFilter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceSpecsPriceFilter struct{}"
	}

	return strings.Join([]string{"ResourceSpecsPriceFilter", string(data)}, " ")
}

type ResourceSpecsPriceFilterKey struct {
	value string
}

type ResourceSpecsPriceFilterKeyEnum struct {
	RESOURCE_SPEC ResourceSpecsPriceFilterKey
	CHARGING_MODE ResourceSpecsPriceFilterKey
}

func GetResourceSpecsPriceFilterKeyEnum() ResourceSpecsPriceFilterKeyEnum {
	return ResourceSpecsPriceFilterKeyEnum{
		RESOURCE_SPEC: ResourceSpecsPriceFilterKey{
			value: "RESOURCE_SPEC",
		},
		CHARGING_MODE: ResourceSpecsPriceFilterKey{
			value: "CHARGING_MODE",
		},
	}
}

func (c ResourceSpecsPriceFilterKey) Value() string {
	return c.value
}

func (c ResourceSpecsPriceFilterKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceSpecsPriceFilterKey) UnmarshalJSON(b []byte) error {
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
