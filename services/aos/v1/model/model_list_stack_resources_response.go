package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStackResourcesResponse Response Object
type ListStackResourcesResponse struct {

	// 资源栈中所管理的资源信息列表
	StackResources *[]StackResource `json:"stack_resources,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListStackResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStackResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListStackResourcesResponse", string(data)}, " ")
}
