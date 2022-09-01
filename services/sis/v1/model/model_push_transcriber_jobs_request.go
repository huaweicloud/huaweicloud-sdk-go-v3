package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type PushTranscriberJobsRequest struct {
	Body *PostTranscriberJobs `json:"body,omitempty" xml:"body"`
}

func (o PushTranscriberJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PushTranscriberJobsRequest struct{}"
	}

	return strings.Join([]string{"PushTranscriberJobsRequest", string(data)}, " ")
}
