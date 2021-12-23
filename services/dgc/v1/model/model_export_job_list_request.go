package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExportJobListRequest struct {
	Body *ExportJobsReq `json:"body,omitempty"`
}

func (o ExportJobListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportJobListRequest struct{}"
	}

	return strings.Join([]string{"ExportJobListRequest", string(data)}, " ")
}
