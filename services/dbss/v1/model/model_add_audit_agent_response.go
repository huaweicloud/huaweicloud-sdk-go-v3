package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddAuditAgentResponse Response Object
type AddAuditAgentResponse struct {

	// 响应状态
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddAuditAgentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAuditAgentResponse struct{}"
	}

	return strings.Join([]string{"AddAuditAgentResponse", string(data)}, " ")
}
