package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDbCacheRuleRequest Request Object
type DeleteDbCacheRuleRequest struct {
	Body *DeleteDbCacheRuleRequestBody `json:"body,omitempty"`
}

func (o DeleteDbCacheRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDbCacheRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteDbCacheRuleRequest", string(data)}, " ")
}
