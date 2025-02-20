package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMultiCloudResourcesResponse Response Object
type ListMultiCloudResourcesResponse struct {

	// 多云资源列表
	Data           *[]BatchListMultiCloudResourceResponseData `json:"data,omitempty"`
	HttpStatusCode int                                        `json:"-"`
}

func (o ListMultiCloudResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMultiCloudResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListMultiCloudResourcesResponse", string(data)}, " ")
}
