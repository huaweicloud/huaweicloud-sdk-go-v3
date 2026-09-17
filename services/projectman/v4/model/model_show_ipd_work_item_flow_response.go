package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIpdWorkItemFlowResponse Response Object
type ShowIpdWorkItemFlowResponse struct {

	// **参数解释**： 状态码。  **取值范围**： 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 响应信息。  **取值范围**： 不涉及。
	Message *string `json:"message,omitempty"`

	Result         *WorkItemFlowInfoVo `json:"result,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowIpdWorkItemFlowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpdWorkItemFlowResponse struct{}"
	}

	return strings.Join([]string{"ShowIpdWorkItemFlowResponse", string(data)}, " ")
}
