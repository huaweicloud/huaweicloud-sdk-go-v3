package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkflowDagPolicies dag policy struct
type WorkflowDagPolicies struct {

	// 是否使用缓存。
	UseCache *bool `json:"use_cache,omitempty"`
}

func (o WorkflowDagPolicies) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowDagPolicies struct{}"
	}

	return strings.Join([]string{"WorkflowDagPolicies", string(data)}, " ")
}
