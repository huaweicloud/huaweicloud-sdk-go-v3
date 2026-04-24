package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOauth2CredentialProvidersRequest Request Object
type ListOauth2CredentialProvidersRequest struct {

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记。
	Marker *string `json:"marker,omitempty"`
}

func (o ListOauth2CredentialProvidersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOauth2CredentialProvidersRequest struct{}"
	}

	return strings.Join([]string{"ListOauth2CredentialProvidersRequest", string(data)}, " ")
}
