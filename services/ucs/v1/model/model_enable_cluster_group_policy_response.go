package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableClusterGroupPolicyResponse Response Object
type EnableClusterGroupPolicyResponse struct {

	// 任务id
	JobID          *string `json:"jobID,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o EnableClusterGroupPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableClusterGroupPolicyResponse struct{}"
	}

	return strings.Join([]string{"EnableClusterGroupPolicyResponse", string(data)}, " ")
}
