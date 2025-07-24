package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunInstancesRequest Request Object
type RunInstancesRequest struct {
	Body *RunInstancesOptions `json:"body,omitempty"`
}

func (o RunInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunInstancesRequest struct{}"
	}

	return strings.Join([]string{"RunInstancesRequest", string(data)}, " ")
}
