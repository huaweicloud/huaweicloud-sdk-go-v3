package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPolicyVersionsV5Response Response Object
type ListPolicyVersionsV5Response struct {

	// 身份策略版本列表。
	Versions *[]PolicyVersion `json:"versions,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListPolicyVersionsV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyVersionsV5Response struct{}"
	}

	return strings.Join([]string{"ListPolicyVersionsV5Response", string(data)}, " ")
}
