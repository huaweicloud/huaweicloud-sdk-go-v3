package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCdnDomainTopRefersResponse Response Object
type ListCdnDomainTopRefersResponse struct {

	// **参数解释：** 数据详情 **取值范围：** 不涉及
	TopReferSummary *[]TopReferSummary `json:"top_refer_summary,omitempty"`
	HttpStatusCode  int                `json:"-"`
}

func (o ListCdnDomainTopRefersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCdnDomainTopRefersResponse struct{}"
	}

	return strings.Join([]string{"ListCdnDomainTopRefersResponse", string(data)}, " ")
}
