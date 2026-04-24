package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStsCredentialProvidersRequest Request Object
type ListStsCredentialProvidersRequest struct {

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记。
	Marker *string `json:"marker,omitempty"`
}

func (o ListStsCredentialProvidersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStsCredentialProvidersRequest struct{}"
	}

	return strings.Join([]string{"ListStsCredentialProvidersRequest", string(data)}, " ")
}
