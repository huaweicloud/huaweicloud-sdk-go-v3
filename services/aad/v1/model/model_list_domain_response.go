package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainResponse Response Object
type ListDomainResponse struct {

	// 域名总数
	Count *int32 `json:"count,omitempty"`

	// 域名列表
	Items          *[]DomainInfo `json:"items,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainResponse struct{}"
	}

	return strings.Join([]string{"ListDomainResponse", string(data)}, " ")
}
