package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ResourceResult struct {

	// 指定类型的配额。 - instance: 表示实例的配额
	Type ResourceResultType `json:"type"`

	// 已创建的资源个数
	Used int32 `json:"used"`

	// 资源最大的配额数
	Quota int32 `json:"quota"`
}

func (o ResourceResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceResult struct{}"
	}

	return strings.Join([]string{"ResourceResult", string(data)}, " ")
}

type ResourceResultType struct {
	value string
}

type ResourceResultTypeEnum struct {
	INSTANCE ResourceResultType
}

func GetResourceResultTypeEnum() ResourceResultTypeEnum {
	return ResourceResultTypeEnum{
		INSTANCE: ResourceResultType{
			value: "instance",
		},
	}
}

func (c ResourceResultType) Value() string {
	return c.value
}

func (c ResourceResultType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceResultType) UnmarshalJSON(b []byte) error {
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
