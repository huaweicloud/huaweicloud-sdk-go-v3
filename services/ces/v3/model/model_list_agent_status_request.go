package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAgentStatusRequest Request Object
type ListAgentStatusRequest struct {
	Body *ListAgentStatusRequestBody `json:"body,omitempty"`
}

func (o ListAgentStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgentStatusRequest struct{}"
	}

	return strings.Join([]string{"ListAgentStatusRequest", string(data)}, " ")
}
