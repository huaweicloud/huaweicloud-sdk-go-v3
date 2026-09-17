package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServiceSpecificCredentialsV5Response Response Object
type ListServiceSpecificCredentialsV5Response struct {

	// 服务专属凭证列表。
	ServiceSpecificCredentials *[]ServiceSpecificCredentialMetadata `json:"service_specific_credentials,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListServiceSpecificCredentialsV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceSpecificCredentialsV5Response struct{}"
	}

	return strings.Join([]string{"ListServiceSpecificCredentialsV5Response", string(data)}, " ")
}
