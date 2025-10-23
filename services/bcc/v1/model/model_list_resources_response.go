package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourcesResponse Response Object
type ListResourcesResponse struct {

	// 资源列表
	Resources *[]ResourceItem `json:"resources,omitempty"`

	// 资源总数
	Count *int64 `json:"count,omitempty"`

	// 下一页的marker
	NextMarker     *string `json:"next_marker,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListResourcesResponse", string(data)}, " ")
}
