package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAuditAgentNewResponse Response Object
type DeleteAuditAgentNewResponse struct {

	// 操作结果  - SUCCESS: 成功
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAuditAgentNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAuditAgentNewResponse struct{}"
	}

	return strings.Join([]string{"DeleteAuditAgentNewResponse", string(data)}, " ")
}
