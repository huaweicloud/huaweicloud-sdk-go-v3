package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCdnDomainTopUasResponse Response Object
type ListCdnDomainTopUasResponse struct {

	// **参数解释：** 数据详情 **取值范围：** 不涉及
	TopUaSummary   *[]TopUaSummary `json:"top_ua_summary,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListCdnDomainTopUasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCdnDomainTopUasResponse struct{}"
	}

	return strings.Join([]string{"ListCdnDomainTopUasResponse", string(data)}, " ")
}
