package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDevServerJobsRequest Request Object
type DeleteDevServerJobsRequest struct {
	Body *ServerJobDeleteRequest `json:"body,omitempty"`
}

func (o DeleteDevServerJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDevServerJobsRequest struct{}"
	}

	return strings.Join([]string{"DeleteDevServerJobsRequest", string(data)}, " ")
}
