package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunRemediationExecutionResponse Response Object
type RunRemediationExecutionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RunRemediationExecutionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunRemediationExecutionResponse struct{}"
	}

	return strings.Join([]string{"RunRemediationExecutionResponse", string(data)}, " ")
}
