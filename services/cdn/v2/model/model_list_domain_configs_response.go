package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainConfigsResponse Response Object
type ListDomainConfigsResponse struct {

	// 域名cname状态。
	CnameStatus    *[]CnameStatus `json:"cname_status,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListDomainConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainConfigsResponse struct{}"
	}

	return strings.Join([]string{"ListDomainConfigsResponse", string(data)}, " ")
}
