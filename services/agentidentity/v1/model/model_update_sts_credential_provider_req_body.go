package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateStsCredentialProviderReqBody struct {

	// The URN of the agency used to obtain IAM temporary credentials.
	AgencyUrn *string `json:"agency_urn,omitempty"`

	// 自定义标签列表。
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o UpdateStsCredentialProviderReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStsCredentialProviderReqBody struct{}"
	}

	return strings.Join([]string{"UpdateStsCredentialProviderReqBody", string(data)}, " ")
}
