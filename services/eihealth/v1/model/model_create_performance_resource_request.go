package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreatePerformanceResourceRequest struct {
	Body *CreatePerformanceResourceReq `json:"body,omitempty"`
}

func (o CreatePerformanceResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePerformanceResourceRequest struct{}"
	}

	return strings.Join([]string{"CreatePerformanceResourceRequest", string(data)}, " ")
}
