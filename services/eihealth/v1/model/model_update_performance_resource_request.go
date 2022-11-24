package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdatePerformanceResourceRequest struct {

	// 性能加速ID
	Id string `json:"id"`

	Body *UpdatePerformanceResourceReq `json:"body,omitempty"`
}

func (o UpdatePerformanceResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePerformanceResourceRequest struct{}"
	}

	return strings.Join([]string{"UpdatePerformanceResourceRequest", string(data)}, " ")
}
