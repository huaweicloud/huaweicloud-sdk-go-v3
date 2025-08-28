package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeLoadbalancerChargeModeResponse Response Object
type ChangeLoadbalancerChargeModeResponse struct {

	// **参数解释**：转包周期下单成功的EIP ID列表。  **取值范围**：不涉及
	EipIdList *[]string `json:"eip_id_list,omitempty"`

	// **参数解释**：转包周期下单成功的LB ID列表。  **取值范围**：不涉及
	LoadbalancerIdList *[]string `json:"loadbalancer_id_list,omitempty"`

	// **参数解释**：转包周期订单号。  **取值范围**：不涉及
	OrderId *string `json:"order_id,omitempty"`

	// **参数解释**：请求ID。  **取值范围**：由数字、小写字母和中划线（-）组成的字符串，自动生成。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeLoadbalancerChargeModeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeLoadbalancerChargeModeResponse struct{}"
	}

	return strings.Join([]string{"ChangeLoadbalancerChargeModeResponse", string(data)}, " ")
}
