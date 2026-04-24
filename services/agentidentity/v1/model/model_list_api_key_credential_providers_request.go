package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApiKeyCredentialProvidersRequest Request Object
type ListApiKeyCredentialProvidersRequest struct {

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记。
	Marker *string `json:"marker,omitempty"`
}

func (o ListApiKeyCredentialProvidersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiKeyCredentialProvidersRequest struct{}"
	}

	return strings.Join([]string{"ListApiKeyCredentialProvidersRequest", string(data)}, " ")
}
