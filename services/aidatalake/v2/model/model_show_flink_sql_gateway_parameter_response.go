package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowFlinkSqlGatewayParameterResponse **参数解释**：SQL Gateway作业参数响应，包含Gateway连接信息。
type ShowFlinkSqlGatewayParameterResponse struct {

	// **参数解释**：作业类型。 **约束限制**：不涉及。 **取值范围**：   - FLINK_SQL_GATEWAY：SQL Gateway作业。
	Type *ShowFlinkSqlGatewayParameterResponseType `json:"type,omitempty"`

	SqlGatewayConfig *ShowFlinkSqlGatewayConfigResponseBody `json:"sql_gateway_config,omitempty"`

	// **参数解释**：是否开启空闲自动释放。 **取值范围**：   - true：开启空闲自动释放。   - false：关闭空闲自动释放。
	IdleEnabled *bool `json:"idle_enabled,omitempty"`

	// **参数解释**：最大空闲时长，单位为分钟。空闲指会话内没有正在运行的作业，创建会话等其他操作不计为非空闲。 **取值范围**：10~1440。
	IdleTimeout *int32 `json:"idle_timeout,omitempty"`
}

func (o ShowFlinkSqlGatewayParameterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlinkSqlGatewayParameterResponse struct{}"
	}

	return strings.Join([]string{"ShowFlinkSqlGatewayParameterResponse", string(data)}, " ")
}

type ShowFlinkSqlGatewayParameterResponseType struct {
	value string
}

type ShowFlinkSqlGatewayParameterResponseTypeEnum struct {
	FLINK_SQL_GATEWAY ShowFlinkSqlGatewayParameterResponseType
}

func GetShowFlinkSqlGatewayParameterResponseTypeEnum() ShowFlinkSqlGatewayParameterResponseTypeEnum {
	return ShowFlinkSqlGatewayParameterResponseTypeEnum{
		FLINK_SQL_GATEWAY: ShowFlinkSqlGatewayParameterResponseType{
			value: "FLINK_SQL_GATEWAY",
		},
	}
}

func (c ShowFlinkSqlGatewayParameterResponseType) Value() string {
	return c.value
}

func (c ShowFlinkSqlGatewayParameterResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFlinkSqlGatewayParameterResponseType) UnmarshalJSON(b []byte) error {
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
