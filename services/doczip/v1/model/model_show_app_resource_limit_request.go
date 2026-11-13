package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAppResourceLimitRequest Request Object
type ShowAppResourceLimitRequest struct {

	// **参数解释**： 查询类型。 **约束限制**： 不涉及。 **取值范围**：   1：获取当前允许创建的应用数量。 **默认取值**： 1。
	Type *ShowAppResourceLimitRequestType `json:"type,omitempty"`

	// **参数解释**： 租户ID。 **约束限制**： 不涉及。 **取值范围**： [1-50]个字符。 **默认取值**： 不涉及。
	TenantId string `json:"tenant_id"`
}

func (o ShowAppResourceLimitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppResourceLimitRequest struct{}"
	}

	return strings.Join([]string{"ShowAppResourceLimitRequest", string(data)}, " ")
}

type ShowAppResourceLimitRequestType struct {
	value int32
}

type ShowAppResourceLimitRequestTypeEnum struct {
	E_1 ShowAppResourceLimitRequestType
}

func GetShowAppResourceLimitRequestTypeEnum() ShowAppResourceLimitRequestTypeEnum {
	return ShowAppResourceLimitRequestTypeEnum{
		E_1: ShowAppResourceLimitRequestType{
			value: 1,
		},
	}
}

func (c ShowAppResourceLimitRequestType) Value() int32 {
	return c.value
}

func (c ShowAppResourceLimitRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAppResourceLimitRequestType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
