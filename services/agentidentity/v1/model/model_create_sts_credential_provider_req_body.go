package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateStsCredentialProviderReqBody struct {

	// The name of the credential provider.
	Name string `json:"name"`

	// The URN of the agency used to obtain IAM temporary credentials.
	AgencyUrn string `json:"agency_urn"`

	// 自定义标签列表。
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o CreateStsCredentialProviderReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStsCredentialProviderReqBody struct{}"
	}

	return strings.Join([]string{"CreateStsCredentialProviderReqBody", string(data)}, " ")
}
