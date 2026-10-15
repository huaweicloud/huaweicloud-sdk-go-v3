package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FlinkSqlGatewayParameter **参数解释**：SQL Gateway作业参数，用于创建交互式SQL Gateway实例。 **约束限制**：不涉及。
type FlinkSqlGatewayParameter struct {

	// **参数解释**：作业类型。 **约束限制**：固定值为 FLINK_SQL_GATEWAY。 **取值范围**：不涉及。 **默认取值**：FLINK_SQL_GATEWAY。
	Type *FlinkSqlGatewayParameterType `json:"type,omitempty"`

	// **参数解释**：用户自定义的SQL Gateway访问鉴权Token，调用方通过HTTP请求头 Authorization: Bearer {token} 携带。 **约束限制**：不涉及。 **取值范围**：长度为1~256个字符。 **默认取值**：不涉及。
	SqlGatewayToken *string `json:"sql_gateway_token,omitempty"`
}

func (o FlinkSqlGatewayParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkSqlGatewayParameter struct{}"
	}

	return strings.Join([]string{"FlinkSqlGatewayParameter", string(data)}, " ")
}

type FlinkSqlGatewayParameterType struct {
	value string
}

type FlinkSqlGatewayParameterTypeEnum struct {
	FLINK_SQL_GATEWAY FlinkSqlGatewayParameterType
}

func GetFlinkSqlGatewayParameterTypeEnum() FlinkSqlGatewayParameterTypeEnum {
	return FlinkSqlGatewayParameterTypeEnum{
		FLINK_SQL_GATEWAY: FlinkSqlGatewayParameterType{
			value: "FLINK_SQL_GATEWAY",
		},
	}
}

func (c FlinkSqlGatewayParameterType) Value() string {
	return c.value
}

func (c FlinkSqlGatewayParameterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlinkSqlGatewayParameterType) UnmarshalJSON(b []byte) error {
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
