package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckKernelUpgradeResponse Response Object
type CheckKernelUpgradeResponse struct {

	// **参数解释**：  预检查任务ID集合。  **取值范围**：  不涉及。
	JobIds         *[]string `json:"job_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CheckKernelUpgradeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckKernelUpgradeResponse struct{}"
	}

	return strings.Join([]string{"CheckKernelUpgradeResponse", string(data)}, " ")
}
