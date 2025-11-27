package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableClusterPolicyResponse Response Object
type DisableClusterPolicyResponse struct {

	// 任务id
	JobID          *string `json:"jobID,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DisableClusterPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableClusterPolicyResponse struct{}"
	}

	return strings.Join([]string{"DisableClusterPolicyResponse", string(data)}, " ")
}
