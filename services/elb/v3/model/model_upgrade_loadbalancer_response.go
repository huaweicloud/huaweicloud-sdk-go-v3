package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeLoadbalancerResponse Response Object
type UpgradeLoadbalancerResponse struct {

	// **参数解释**：请求ID。  **取值范围**：由数字、小写字母和中划线（-）组成的字符串，自动生成。
	RequestId *string `json:"request_id,omitempty"`

	// **参数解释**：升级的任务ID。  **取值范围**：不涉及
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpgradeLoadbalancerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeLoadbalancerResponse struct{}"
	}

	return strings.Join([]string{"UpgradeLoadbalancerResponse", string(data)}, " ")
}
