package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobResourcesResponse Response Object
type ListJobResourcesResponse struct {

	// 资源包返回总数
	Total *int32 `json:"total,omitempty"`

	// 已上传的用户资源名列表。
	Resources *[]PackageResource `json:"resources,omitempty"`

	// 系统内置资源模块列表
	Modules *[]PackageResourceMoudle `json:"modules,omitempty"`

	// 已上传的用户分组资源。
	Groups         *[]PackageResourceGroup `json:"groups,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListJobResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListJobResourcesResponse", string(data)}, " ")
}
