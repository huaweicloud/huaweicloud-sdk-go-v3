package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateWorkloadIdentityReqBody struct {

	// The name of the workload identity.
	Name string `json:"name"`

	AllowedResourceOauth2ReturnUrls *[]string `json:"allowed_resource_oauth2_return_urls,omitempty"`

	AuthorizerType *AuthorizerType `json:"authorizer_type"`

	AuthorizerConfiguration *AuthorizerConfiguration `json:"authorizer_configuration,omitempty"`

	// 自定义标签列表。
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o CreateWorkloadIdentityReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkloadIdentityReqBody struct{}"
	}

	return strings.Join([]string{"CreateWorkloadIdentityReqBody", string(data)}, " ")
}
