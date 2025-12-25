package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AuditDo struct {

	// **参数解释**： 操作ID。  **取值范围**： 不涉及。id
	Id *string `json:"id,omitempty"`

	// **参数解释**： 操作模块。  **取值范围**： 不涉及。
	Module *string `json:"module,omitempty"`

	// **参数解释**： 操作类型。  **取值范围**： 不涉及。
	Operation *string `json:"operation,omitempty"`

	// **参数解释**： 操作时间。  **取值范围**： 不涉及。
	Time *string `json:"time,omitempty"`

	// **参数解释**： 操作信息。  **取值范围**： 不涉及。
	Info *string `json:"info,omitempty"`

	// **参数解释**： 操作数据。  **取值范围**： 不涉及。
	Data *string `json:"data,omitempty"`

	// **参数解释**： 原因。  **取值范围**： 不涉及。
	Reason *string `json:"reason,omitempty"`

	// **参数解释**： 操作人客户端类型。  **取值范围**： 不涉及。
	UserAgent *string `json:"userAgent,omitempty"`

	// **参数解释**： 操作人。  **取值范围**： 不涉及。
	Operator *string `json:"operator,omitempty"`

	// **参数解释**： 操作人名字。  **取值范围**： 不涉及。
	OperatorNickName *string `json:"operatorNickName,omitempty"`

	// **参数解释**： 操作人租户名。  **取值范围**： 不涉及。
	OperatorTenantName *string `json:"operatorTenantName,omitempty"`

	// **参数解释**： 操作人IP地址。  **取值范围**： 不涉及。
	IpSource *string `json:"ipSource,omitempty"`
}

func (o AuditDo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuditDo struct{}"
	}

	return strings.Join([]string{"AuditDo", string(data)}, " ")
}
