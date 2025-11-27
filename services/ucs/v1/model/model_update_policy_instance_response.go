package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePolicyInstanceResponse Response Object
type UpdatePolicyInstanceResponse struct {

	// 任务id
	JobID          *string `json:"jobID,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePolicyInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyInstanceResponse struct{}"
	}

	return strings.Join([]string{"UpdatePolicyInstanceResponse", string(data)}, " ")
}
