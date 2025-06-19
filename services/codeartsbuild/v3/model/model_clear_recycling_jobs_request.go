package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClearRecyclingJobsRequest Request Object
type ClearRecyclingJobsRequest struct {
}

func (o ClearRecyclingJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClearRecyclingJobsRequest struct{}"
	}

	return strings.Join([]string{"ClearRecyclingJobsRequest", string(data)}, " ")
}
