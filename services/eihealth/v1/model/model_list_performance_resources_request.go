package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPerformanceResourcesRequest Request Object
type ListPerformanceResourcesRequest struct {
}

func (o ListPerformanceResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPerformanceResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListPerformanceResourcesRequest", string(data)}, " ")
}
