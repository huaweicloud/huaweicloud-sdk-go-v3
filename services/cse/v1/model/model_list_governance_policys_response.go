package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGovernancePolicysResponse Response Object
type ListGovernancePolicysResponse struct {

	// 查询治理策略列表响应结构体
	Body           *[]ListGovernancePolicyResponseBody `json:"body,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ListGovernancePolicysResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGovernancePolicysResponse struct{}"
	}

	return strings.Join([]string{"ListGovernancePolicysResponse", string(data)}, " ")
}
