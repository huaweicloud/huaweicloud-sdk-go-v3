package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePolicyInstanceResponse Response Object
type DeletePolicyInstanceResponse struct {

	// 任务id
	JobID          *string `json:"jobID,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeletePolicyInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePolicyInstanceResponse struct{}"
	}

	return strings.Join([]string{"DeletePolicyInstanceResponse", string(data)}, " ")
}
