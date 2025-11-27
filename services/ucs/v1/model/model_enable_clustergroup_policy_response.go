package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableClustergroupPolicyResponse Response Object
type EnableClustergroupPolicyResponse struct {

	// 任务id
	JobID          *string `json:"jobID,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o EnableClustergroupPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableClustergroupPolicyResponse struct{}"
	}

	return strings.Join([]string{"EnableClustergroupPolicyResponse", string(data)}, " ")
}
