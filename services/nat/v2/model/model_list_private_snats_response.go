package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPrivateSnatsResponse Response Object
type ListPrivateSnatsResponse struct {

	// 查询SNAT规则列表的响应体。
	SnatRules *[]PrivateSnat `json:"snat_rules,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPrivateSnatsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateSnatsResponse struct{}"
	}

	return strings.Join([]string{"ListPrivateSnatsResponse", string(data)}, " ")
}
