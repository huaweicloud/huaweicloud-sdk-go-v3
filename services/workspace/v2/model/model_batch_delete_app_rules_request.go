package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteAppRulesRequest Request Object
type BatchDeleteAppRulesRequest struct {
	Body *BatchOperateAppRulesReq `json:"body,omitempty"`
}

func (o BatchDeleteAppRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAppRulesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteAppRulesRequest", string(data)}, " ")
}
