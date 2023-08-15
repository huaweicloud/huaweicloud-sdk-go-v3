package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourcesResponse Response Object
type ListResourcesResponse struct {

	// 已上传的用户资源名列表。
	Resources *[]ListGroupPackagesResource `json:"resources,omitempty"`

	// 系统内置资源模块列表
	Modules *[]ListResourcePackagesRespMoudle `json:"modules,omitempty"`

	// 已上传的用户分组资源。
	Groups *[]interface{} `json:"groups,omitempty"`

	// 资源包返回总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListResourcesResponse", string(data)}, " ")
}
