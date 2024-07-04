package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomIngressPortDomainsResponse Response Object
type ListCustomIngressPortDomainsResponse struct {

	// 本次返回的列表长度
	Size int32 `json:"size"`

	// 满足条件的记录数
	Total int64 `json:"total"`

	// 入方向端口绑定的域名信息列表。
	DomainInfos    *[]PortBindingDomainInfo `json:"domain_infos,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListCustomIngressPortDomainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomIngressPortDomainsResponse struct{}"
	}

	return strings.Join([]string{"ListCustomIngressPortDomainsResponse", string(data)}, " ")
}
