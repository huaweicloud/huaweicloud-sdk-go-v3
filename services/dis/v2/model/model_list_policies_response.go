package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPoliciesResponse Response Object
type ListPoliciesResponse struct {

	// 通道唯一标识符。
	StreamId *string `json:"stream_id,omitempty"`

	// 通道授权信息列表。
	Rules          *[]PrincipalRule `json:"rules,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListPoliciesResponse", string(data)}, " ")
}
