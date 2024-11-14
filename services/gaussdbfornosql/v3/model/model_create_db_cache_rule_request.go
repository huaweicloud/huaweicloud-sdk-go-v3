package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDbCacheRuleRequest Request Object
type CreateDbCacheRuleRequest struct {
	Body *CreateDbCacheRuleRequestBody `json:"body,omitempty"`
}

func (o CreateDbCacheRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDbCacheRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateDbCacheRuleRequest", string(data)}, " ")
}
