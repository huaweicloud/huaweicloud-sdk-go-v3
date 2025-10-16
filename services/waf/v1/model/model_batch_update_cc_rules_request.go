package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateCcRulesRequest Request Object
type BatchUpdateCcRulesRequest struct {
	Body *BatchUpdateCcRulesRequestBody `json:"body,omitempty"`
}

func (o BatchUpdateCcRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateCcRulesRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateCcRulesRequest", string(data)}, " ")
}
