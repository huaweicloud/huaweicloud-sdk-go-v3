package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreLoadbalancerResponse Response Object
type RestoreLoadbalancerResponse struct {

	// **参数解释**：请求ID。  **取值范围**：由数字、小写字母和中划线（-）组成的字符串，自动生成。
	RequestId *string `json:"request_id,omitempty"`

	// **参数解释**：ELB实例的类型。  **取值范围**： - v2：共享型ELB。 - v3：独享型ELB。
	Type *string `json:"type,omitempty"`

	// **参数解释**：还原负载均衡器的任务ID。  **取值范围**：不涉及
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestoreLoadbalancerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreLoadbalancerResponse struct{}"
	}

	return strings.Join([]string{"RestoreLoadbalancerResponse", string(data)}, " ")
}
