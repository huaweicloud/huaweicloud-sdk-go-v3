package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlinkSqlGatewayConfigResponseBody **参数解释**：SQL Gateway连接信息。
type ShowFlinkSqlGatewayConfigResponseBody struct {

	// **参数解释**：SQL Gateway的Ingress访问域名，格式为sql-gateway-{jobId}。 **取值范围**：长度为1~253个字符。
	Host *string `json:"host,omitempty"`

	// **参数解释**：SQL Gateway鉴权Token。 **取值范围**：已配置时固定返回 ******；未配置时不返回该字段。
	SqlGatewayToken *string `json:"sql_gateway_token,omitempty"`
}

func (o ShowFlinkSqlGatewayConfigResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlinkSqlGatewayConfigResponseBody struct{}"
	}

	return strings.Join([]string{"ShowFlinkSqlGatewayConfigResponseBody", string(data)}, " ")
}
