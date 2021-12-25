package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchDeleteRulesRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`

	Body *BatchDeleteRulesRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteRulesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteRulesRequest", string(data)}, " ")
}
