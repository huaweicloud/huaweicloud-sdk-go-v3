package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTopDomainsResponse Response Object
type ListTopDomainsResponse struct {

	// CountItem的总数量
	Total *int32 `json:"total,omitempty"`

	// CountItem详细信息
	Items          *[]TopDomainsCountItem `json:"items,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListTopDomainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopDomainsResponse struct{}"
	}

	return strings.Join([]string{"ListTopDomainsResponse", string(data)}, " ")
}
