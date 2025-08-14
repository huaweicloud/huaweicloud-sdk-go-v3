package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCatalogApplicationsRequest Request Object
type ListCatalogApplicationsRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 每个请求返回的最大结果数。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记
	Marker *string `json:"marker,omitempty"`
}

func (o ListCatalogApplicationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCatalogApplicationsRequest struct{}"
	}

	return strings.Join([]string{"ListCatalogApplicationsRequest", string(data)}, " ")
}
