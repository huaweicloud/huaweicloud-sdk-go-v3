package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGroupSummaryResponse Response Object
type ShowGroupSummaryResponse struct {

	// 用户组列表。
	Groups *[]GroupSummary `json:"groups,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowGroupSummaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupSummaryResponse struct{}"
	}

	return strings.Join([]string{"ShowGroupSummaryResponse", string(data)}, " ")
}
