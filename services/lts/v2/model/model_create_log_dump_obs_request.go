package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateLogDumpObsRequest struct {
	Body *CreateLogDumpObsRequestBody `json:"body,omitempty"`
}

func (o CreateLogDumpObsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogDumpObsRequest struct{}"
	}

	return strings.Join([]string{"CreateLogDumpObsRequest", string(data)}, " ")
}
