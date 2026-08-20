package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTopUrlResponse Response Object
type ShowTopUrlResponse struct {

	// **参数解释：** 服务范围 **取值范围：** - mainland_china：中国大陆 - outside_mainland_china：中国大陆境外 - global：全球
	ServiceArea *string `json:"service_area,omitempty"`

	// **参数解释：** 数据详情 **取值范围：** 不涉及
	TopUrlSummary  *[]TopUrlSummary `json:"top_url_summary,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowTopUrlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTopUrlResponse struct{}"
	}

	return strings.Join([]string{"ShowTopUrlResponse", string(data)}, " ")
}
