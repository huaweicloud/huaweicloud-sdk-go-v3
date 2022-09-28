package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPerformanceResourceStatResponse struct {

	// 性能加速资源总数
	Count *int32 `json:"count,omitempty"`

	// 性能加速资源信息
	PerformanceResources *[]PerformanceResourcesRsp `json:"performance_resources,omitempty"`
	HttpStatusCode       int                        `json:"-"`
}

func (o ListPerformanceResourceStatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPerformanceResourceStatResponse struct{}"
	}

	return strings.Join([]string{"ListPerformanceResourceStatResponse", string(data)}, " ")
}
