package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateWorkloadIdentityReqBody struct {
	AllowedResourceOauth2ReturnUrls *[]string `json:"allowed_resource_oauth2_return_urls,omitempty"`

	AuthorizerConfiguration *AuthorizerConfiguration `json:"authorizer_configuration,omitempty"`

	// 自定义标签列表。
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o UpdateWorkloadIdentityReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkloadIdentityReqBody struct{}"
	}

	return strings.Join([]string{"UpdateWorkloadIdentityReqBody", string(data)}, " ")
}
