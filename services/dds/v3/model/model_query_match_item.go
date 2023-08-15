package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type QueryMatchItem struct {

	// 取值为“instance_name”或“instance_id”，分别表示按实例名称或按实例ID匹配查询。
	Key QueryMatchItemKey `json:"key"`

	// 待匹配的实例名称或实例ID，可以调用“查询实例列表和详情”接口获取。如果未申请实例，可以调用“创建实例”接口创建。
	Value string `json:"value"`
}

func (o QueryMatchItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryMatchItem struct{}"
	}

	return strings.Join([]string{"QueryMatchItem", string(data)}, " ")
}

type QueryMatchItemKey struct {
	value string
}

type QueryMatchItemKeyEnum struct {
	INSTANCE_NAME QueryMatchItemKey
	INSTANCE_ID   QueryMatchItemKey
}

func GetQueryMatchItemKeyEnum() QueryMatchItemKeyEnum {
	return QueryMatchItemKeyEnum{
		INSTANCE_NAME: QueryMatchItemKey{
			value: "instance_name",
		},
		INSTANCE_ID: QueryMatchItemKey{
			value: "instance_id",
		},
	}
}

func (c QueryMatchItemKey) Value() string {
	return c.value
}

func (c QueryMatchItemKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryMatchItemKey) UnmarshalJSON(b []byte) error {
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
