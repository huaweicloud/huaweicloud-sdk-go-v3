package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPolicyResponse Response Object
type ListPolicyResponse struct {
	PageInfo *PagedInfo `json:"page_info,omitempty"`

	// 分页信息
	Policies       *[]Policy `json:"policies,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyResponse struct{}"
	}

	return strings.Join([]string{"ListPolicyResponse", string(data)}, " ")
}
