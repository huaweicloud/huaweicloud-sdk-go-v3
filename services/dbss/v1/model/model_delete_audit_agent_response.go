package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAuditAgentResponse Response Object
type DeleteAuditAgentResponse struct {

	// 响应状态
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAuditAgentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAuditAgentResponse struct{}"
	}

	return strings.Join([]string{"DeleteAuditAgentResponse", string(data)}, " ")
}
