package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlinkSqlGatewayResourceSpecResponseBody **参数解释**：SQL Gateway 的资源规格，包含CPU和内存配额。
type ShowFlinkSqlGatewayResourceSpecResponseBody struct {

	// **参数解释**：CPU毫核数，单位为毫核。 **取值范围**：500~47500。
	Cpu int32 `json:"cpu"`

	// **参数解释**：内存大小，单位为MB。 **取值范围**：2048~391168。
	Memory int32 `json:"memory"`
}

func (o ShowFlinkSqlGatewayResourceSpecResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlinkSqlGatewayResourceSpecResponseBody struct{}"
	}

	return strings.Join([]string{"ShowFlinkSqlGatewayResourceSpecResponseBody", string(data)}, " ")
}
