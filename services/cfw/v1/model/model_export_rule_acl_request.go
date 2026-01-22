package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportRuleAclRequest Request Object
type ExportRuleAclRequest struct {
	Body *ExportRuleAclRequestBody `json:"body,omitempty"`
}

func (o ExportRuleAclRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportRuleAclRequest struct{}"
	}

	return strings.Join([]string{"ExportRuleAclRequest", string(data)}, " ")
}
