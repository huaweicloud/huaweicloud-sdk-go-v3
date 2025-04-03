package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServiceProviderConfigResponse Response Object
type ServiceProviderConfigResponse struct {

	// 概要
	Schemas *[]string `json:"schemas,omitempty"`

	// 帮助文档链接
	DocumentationUri *string `json:"documentationUri,omitempty"`

	// 认证概要列表
	AuthenticationSchemes *[]AuthenticationSchemeItemDto `json:"authenticationSchemes,omitempty"`

	Patch *PatchDto `json:"patch,omitempty"`

	Bulk *BulkDto `json:"bulk,omitempty"`

	Filter *FilterDto `json:"filter,omitempty"`

	ChangePassword *ChangePasswordDto `json:"changePassword,omitempty"`

	Sort *SortDto `json:"sort,omitempty"`

	Etag           *EtagDto `json:"etag,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ServiceProviderConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceProviderConfigResponse struct{}"
	}

	return strings.Join([]string{"ServiceProviderConfigResponse", string(data)}, " ")
}
