package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpgradeEndpointRequestBody 升级终端节点终端节点接口请求结构体
type UpgradeEndpointRequestBody struct {

	// 升级操作。 默认取值为start，表示开始升级。
	Action UpgradeEndpointRequestBodyAction `json:"action"`
}

func (o UpgradeEndpointRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeEndpointRequestBody struct{}"
	}

	return strings.Join([]string{"UpgradeEndpointRequestBody", string(data)}, " ")
}

type UpgradeEndpointRequestBodyAction struct {
	value string
}

type UpgradeEndpointRequestBodyActionEnum struct {
	START UpgradeEndpointRequestBodyAction
}

func GetUpgradeEndpointRequestBodyActionEnum() UpgradeEndpointRequestBodyActionEnum {
	return UpgradeEndpointRequestBodyActionEnum{
		START: UpgradeEndpointRequestBodyAction{
			value: "start",
		},
	}
}

func (c UpgradeEndpointRequestBodyAction) Value() string {
	return c.value
}

func (c UpgradeEndpointRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpgradeEndpointRequestBodyAction) UnmarshalJSON(b []byte) error {
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
