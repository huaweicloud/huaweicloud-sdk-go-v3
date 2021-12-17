package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListComputingResourcesResponse struct {
	// 计算资源总个数。

	Count *int64 `json:"count,omitempty"`
	// 计算资源列表。

	ComputingResources *[]ComputingResource `json:"computing_resources,omitempty"`
	HttpStatusCode     int                  `json:"-"`
}

func (o ListComputingResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListComputingResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListComputingResourcesResponse", string(data)}, " ")
}
