package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunSummaryRequest struct {
	Body *SummaryReq `json:"body,omitempty"`
}

func (o RunSummaryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunSummaryRequest struct{}"
	}

	return strings.Join([]string{"RunSummaryRequest", string(data)}, " ")
}
