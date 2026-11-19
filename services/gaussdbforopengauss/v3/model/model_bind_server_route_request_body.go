package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BindServerRouteRequestBody struct {

	// **参数解释**: 操作类型。 **约束限制**: 不涉及。 **取值范围**: - add：绑定服务端IP路由 - delete：解绑服务端IP路由  **默认取值**: 不涉及。
	Operate BindServerRouteRequestBodyOperate `json:"operate"`

	// **参数解释**: 服务端IP列表。 **约束限制**: 列表元素需要满足IPv4格式要求。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	TargetServerIps []string `json:"target_server_ips"`
}

func (o BindServerRouteRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindServerRouteRequestBody struct{}"
	}

	return strings.Join([]string{"BindServerRouteRequestBody", string(data)}, " ")
}

type BindServerRouteRequestBodyOperate struct {
	value string
}

type BindServerRouteRequestBodyOperateEnum struct {
	ADD    BindServerRouteRequestBodyOperate
	DELETE BindServerRouteRequestBodyOperate
}

func GetBindServerRouteRequestBodyOperateEnum() BindServerRouteRequestBodyOperateEnum {
	return BindServerRouteRequestBodyOperateEnum{
		ADD: BindServerRouteRequestBodyOperate{
			value: "add",
		},
		DELETE: BindServerRouteRequestBodyOperate{
			value: "delete",
		},
	}
}

func (c BindServerRouteRequestBodyOperate) Value() string {
	return c.value
}

func (c BindServerRouteRequestBodyOperate) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BindServerRouteRequestBodyOperate) UnmarshalJSON(b []byte) error {
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
