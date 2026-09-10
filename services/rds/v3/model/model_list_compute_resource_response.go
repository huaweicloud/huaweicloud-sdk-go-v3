package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListComputeResourceResponse Response Object
type ListComputeResourceResponse struct {

	// 资源包信息列表。
	ResourcePackageInfos *[]ResourcePackageInfo `json:"resource_package_infos,omitempty"`

	// 总记录数。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListComputeResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListComputeResourceResponse struct{}"
	}

	return strings.Join([]string{"ListComputeResourceResponse", string(data)}, " ")
}
