package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEventExportJobRequest Request Object
type CreateEventExportJobRequest struct {
	Body *AsyncJobReqBody `json:"body,omitempty"`
}

func (o CreateEventExportJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEventExportJobRequest struct{}"
	}

	return strings.Join([]string{"CreateEventExportJobRequest", string(data)}, " ")
}
