package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCdnDomainTopOriginUrlResponse Response Object
type ListCdnDomainTopOriginUrlResponse struct {

	// 详情数据对象。
	TopOriginUrlSummary *[]TopOriginUrlSummary `json:"top_origin_url_summary,omitempty"`
	HttpStatusCode      int                    `json:"-"`
}

func (o ListCdnDomainTopOriginUrlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCdnDomainTopOriginUrlResponse struct{}"
	}

	return strings.Join([]string{"ListCdnDomainTopOriginUrlResponse", string(data)}, " ")
}
