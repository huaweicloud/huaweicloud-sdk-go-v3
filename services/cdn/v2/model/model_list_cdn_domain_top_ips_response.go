package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCdnDomainTopIpsResponse Response Object
type ListCdnDomainTopIpsResponse struct {

	// 详情数据对象。
	TopIpSummary   *[]TopIpSummary `json:"top_ip_summary,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListCdnDomainTopIpsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCdnDomainTopIpsResponse struct{}"
	}

	return strings.Join([]string{"ListCdnDomainTopIpsResponse", string(data)}, " ")
}
