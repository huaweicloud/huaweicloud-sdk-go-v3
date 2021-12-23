package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListJobInstancesRequest struct {
}

func (o ListJobInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListJobInstancesRequest", string(data)}, " ")
}
