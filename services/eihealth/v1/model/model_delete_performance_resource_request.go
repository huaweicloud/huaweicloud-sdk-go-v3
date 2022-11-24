package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeletePerformanceResourceRequest struct {

	// 性能加速ID
	Id string `json:"id"`
}

func (o DeletePerformanceResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePerformanceResourceRequest struct{}"
	}

	return strings.Join([]string{"DeletePerformanceResourceRequest", string(data)}, " ")
}
