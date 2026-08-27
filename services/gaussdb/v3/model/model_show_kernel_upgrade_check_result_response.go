package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKernelUpgradeCheckResultResponse Response Object
type ShowKernelUpgradeCheckResultResponse struct {

	// **参数解释**：  预检查任务ID集合。  **取值范围**：  不涉及。
	JobIds         *[]string `json:"job_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowKernelUpgradeCheckResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKernelUpgradeCheckResultResponse struct{}"
	}

	return strings.Join([]string{"ShowKernelUpgradeCheckResultResponse", string(data)}, " ")
}
