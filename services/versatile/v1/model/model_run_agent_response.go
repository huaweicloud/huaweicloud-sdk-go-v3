package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunAgentResponse Response Object
type RunAgentResponse struct {

	// 会话id
	ConversationId *string `json:"conversation_id,omitempty"`

	Outputs map[string]interface{} `json:"outputs,omitempty"`

	Messages *[]NodeMessage `json:"messages,omitempty"`

	// 会话id
	ErrorCode *string `json:"error_code,omitempty"`

	// 会话id
	ErrorMessage *string `json:"error_message,omitempty"`

	Metadata map[string]interface{} `json:"metadata,omitempty"`

	Status *Status `json:"status,omitempty"`

	// 开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 结束时间
	EndTime *int64 `json:"end_time,omitempty"`

	NodeInfo *[]NodeRunInfo `json:"node_info,omitempty"`

	Events         *[]WorkflowRunStreamRsp `json:"events,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o RunAgentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunAgentResponse struct{}"
	}

	return strings.Join([]string{"RunAgentResponse", string(data)}, " ")
}
