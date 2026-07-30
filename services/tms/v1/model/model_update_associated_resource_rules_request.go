package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAssociatedResourceRulesRequest Request Object
type UpdateAssociatedResourceRulesRequest struct {
	Body *ReqUpdateAssociatedResourceRules `json:"body,omitempty"`
}

func (o UpdateAssociatedResourceRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAssociatedResourceRulesRequest struct{}"
	}

	return strings.Join([]string{"UpdateAssociatedResourceRulesRequest", string(data)}, " ")
}
