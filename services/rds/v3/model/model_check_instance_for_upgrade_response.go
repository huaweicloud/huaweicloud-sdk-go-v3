package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckInstanceForUpgradeResponse Response Object
type CheckInstanceForUpgradeResponse struct {

	// 工作流id
	WorkflowId *string `json:"workflow_id,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckInstanceForUpgradeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckInstanceForUpgradeResponse struct{}"
	}

	return strings.Join([]string{"CheckInstanceForUpgradeResponse", string(data)}, " ")
}
