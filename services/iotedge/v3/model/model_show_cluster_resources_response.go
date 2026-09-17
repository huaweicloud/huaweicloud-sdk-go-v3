package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterResourcesResponse Response Object
type ShowClusterResourcesResponse struct {

	// 总记录数
	Count *int64 `json:"count,omitempty"`

	PageInfo *PageInfoDto `json:"page_info,omitempty"`

	// 查询资源列表返回消息体
	ResourceList   *[]ResourceDetail `json:"resource_list,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowClusterResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterResourcesResponse struct{}"
	}

	return strings.Join([]string{"ShowClusterResourcesResponse", string(data)}, " ")
}
