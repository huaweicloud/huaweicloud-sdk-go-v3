package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddAuditAgentNewResponse Response Object
type AddAuditAgentNewResponse struct {

	// 操作结果  - SUCCESS：成功
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddAuditAgentNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAuditAgentNewResponse struct{}"
	}

	return strings.Join([]string{"AddAuditAgentNewResponse", string(data)}, " ")
}
