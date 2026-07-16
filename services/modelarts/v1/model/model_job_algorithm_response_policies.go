package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobAlgorithmResponsePolicies 作业支持的策略。
type JobAlgorithmResponsePolicies struct {
	AutoSearch *JobAlgorithmResponsePoliciesAutoSearch `json:"auto_search,omitempty"`
}

func (o JobAlgorithmResponsePolicies) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobAlgorithmResponsePolicies struct{}"
	}

	return strings.Join([]string{"JobAlgorithmResponsePolicies", string(data)}, " ")
}
