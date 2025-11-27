package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableClustergroupPolicyResponse Response Object
type DisableClustergroupPolicyResponse struct {

	// 任务id
	JobID          *string `json:"jobID,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DisableClustergroupPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableClustergroupPolicyResponse struct{}"
	}

	return strings.Join([]string{"DisableClustergroupPolicyResponse", string(data)}, " ")
}
