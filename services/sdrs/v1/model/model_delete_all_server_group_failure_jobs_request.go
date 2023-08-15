package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAllServerGroupFailureJobsRequest Request Object
type DeleteAllServerGroupFailureJobsRequest struct {
}

func (o DeleteAllServerGroupFailureJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAllServerGroupFailureJobsRequest struct{}"
	}

	return strings.Join([]string{"DeleteAllServerGroupFailureJobsRequest", string(data)}, " ")
}
