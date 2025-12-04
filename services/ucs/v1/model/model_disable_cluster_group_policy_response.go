package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableClusterGroupPolicyResponse Response Object
type DisableClusterGroupPolicyResponse struct {

	// 任务id
	JobID          *string `json:"jobID,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DisableClusterGroupPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableClusterGroupPolicyResponse struct{}"
	}

	return strings.Join([]string{"DisableClusterGroupPolicyResponse", string(data)}, " ")
}
