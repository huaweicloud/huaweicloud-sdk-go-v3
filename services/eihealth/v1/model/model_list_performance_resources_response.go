package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPerformanceResourcesResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 性能加速资源列表
	Resources      *[]PerformanceResourceRsp `json:"resources,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListPerformanceResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPerformanceResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListPerformanceResourcesResponse", string(data)}, " ")
}
