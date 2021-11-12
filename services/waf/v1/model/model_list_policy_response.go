package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPolicyResponse struct {
	// 防护策略的数量

	Total *int32 `json:"total,omitempty"`
	// 防护策略的具体内容

	Items          *[]PolicyResponse `json:"items,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyResponse struct{}"
	}

	return strings.Join([]string{"ListPolicyResponse", string(data)}, " ")
}
