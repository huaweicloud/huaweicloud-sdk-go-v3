package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAssociatedResourceRulesRequest Request Object
type CreateAssociatedResourceRulesRequest struct {
	Body *ReqCreateAssociatedResourceRules `json:"body,omitempty"`
}

func (o CreateAssociatedResourceRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAssociatedResourceRulesRequest struct{}"
	}

	return strings.Join([]string{"CreateAssociatedResourceRulesRequest", string(data)}, " ")
}
