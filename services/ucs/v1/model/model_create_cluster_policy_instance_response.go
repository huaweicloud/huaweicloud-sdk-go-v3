package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusterPolicyInstanceResponse Response Object
type CreateClusterPolicyInstanceResponse struct {

	// 任务id
	JobID          *string `json:"jobID,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateClusterPolicyInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterPolicyInstanceResponse struct{}"
	}

	return strings.Join([]string{"CreateClusterPolicyInstanceResponse", string(data)}, " ")
}
