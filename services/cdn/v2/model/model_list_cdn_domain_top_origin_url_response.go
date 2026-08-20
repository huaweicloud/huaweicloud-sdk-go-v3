package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCdnDomainTopOriginUrlResponse Response Object
type ListCdnDomainTopOriginUrlResponse struct {

	// **参数解释：** 数据详情 **取值范围：** 不涉及
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
