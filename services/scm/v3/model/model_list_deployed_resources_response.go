package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDeployedResourcesResponse Response Object
type ListDeployedResourcesResponse struct {

	// 请求结果列表，详情请参见ResultDetail字段数据结构说明。
	Results        *[]ResultDetail `json:"results,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListDeployedResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDeployedResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListDeployedResourcesResponse", string(data)}, " ")
}
