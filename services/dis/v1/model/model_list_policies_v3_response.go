package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPoliciesV3Response struct {
	// 通道唯一标识符。

	StreamId *string `json:"stream_id,omitempty"`
	// 通道授权信息列表。

	Rules          *[]PrincipalRule `json:"rules,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListPoliciesV3Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPoliciesV3Response struct{}"
	}

	return strings.Join([]string{"ListPoliciesV3Response", string(data)}, " ")
}
