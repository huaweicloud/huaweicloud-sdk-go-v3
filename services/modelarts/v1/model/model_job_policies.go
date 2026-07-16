package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobPolicies 作业支持的策略，用于超参搜索。
type JobPolicies struct {
	AutoSearch *AutoSearch `json:"auto_search,omitempty"`
}

func (o JobPolicies) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobPolicies struct{}"
	}

	return strings.Join([]string{"JobPolicies", string(data)}, " ")
}
