package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAgentInfosRequest Request Object
type ShowAgentInfosRequest struct {
	Body *AgentInfoParam `json:"body,omitempty"`
}

func (o ShowAgentInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgentInfosRequest struct{}"
	}

	return strings.Join([]string{"ShowAgentInfosRequest", string(data)}, " ")
}
