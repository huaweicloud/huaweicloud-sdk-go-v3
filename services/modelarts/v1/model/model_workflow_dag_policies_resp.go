package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkflowDagPoliciesResp dag policy struct
type WorkflowDagPoliciesResp struct {

	// **参数解释**：是否使用缓存。 **取值范围**： - true：使用缓存 - false：不使用缓存
	UseCache *bool `json:"use_cache,omitempty"`
}

func (o WorkflowDagPoliciesResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowDagPoliciesResp struct{}"
	}

	return strings.Join([]string{"WorkflowDagPoliciesResp", string(data)}, " ")
}
