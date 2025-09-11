package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAuditDbAgentResponse Response Object
type CreateAuditDbAgentResponse struct {

	// 操作结果  - SUCCESS: 成功
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAuditDbAgentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAuditDbAgentResponse struct{}"
	}

	return strings.Join([]string{"CreateAuditDbAgentResponse", string(data)}, " ")
}
