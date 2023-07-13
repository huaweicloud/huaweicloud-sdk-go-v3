package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowApplicationDependentResourcesResponse Response Object
type ShowApplicationDependentResourcesResponse struct {

	// 依赖云资源信息
	DependentServices *[]ResouceInfo `json:"dependent_services,omitempty"`

	// 资源总个数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowApplicationDependentResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApplicationDependentResourcesResponse struct{}"
	}

	return strings.Join([]string{"ShowApplicationDependentResourcesResponse", string(data)}, " ")
}
