package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDelegatedAdministratorsRequest Request Object
type ListDelegatedAdministratorsRequest struct {

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
