package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAttachedDomainsV2Response struct {

	// 本次返回的列表长度
	Size int32 `json:"size"`

	// 满足条件的记录数
	Total int64 `json:"total"`

	// 已绑定域名集合
	BoundDomains   *[]UrlDomainRefInfo `json:"bound_domains,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListAttachedDomainsV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAttachedDomainsV2Response struct{}"
	}

	return strings.Join([]string{"ListAttachedDomainsV2Response", string(data)}, " ")
}
