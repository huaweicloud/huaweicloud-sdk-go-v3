package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableClusterPolicyResponse Response Object
type EnableClusterPolicyResponse struct {

	// 任务id
	JobID          *string `json:"jobID,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o EnableClusterPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableClusterPolicyResponse struct{}"
	}

	return strings.Join([]string{"EnableClusterPolicyResponse", string(data)}, " ")
}
