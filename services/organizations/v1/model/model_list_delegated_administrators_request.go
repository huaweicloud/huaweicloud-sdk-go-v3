package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDelegatedAdministratorsRequest Request Object
type ListDelegatedAdministratorsRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 服务主体的名称。
	ServicePrincipal *string `json:"service_principal,omitempty"`

	// 页面中最大结果数量。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记。
	Marker *string `json:"marker,omitempty"`
}

func (o ListDelegatedAdministratorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDelegatedAdministratorsRequest struct{}"
	}

	return strings.Join([]string{"ListDelegatedAdministratorsRequest", string(data)}, " ")
}
