package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResourcesSummaryResponse Response Object
type ShowResourcesSummaryResponse struct {

	// 企业项目ID
	EpId *string `json:"ep_id,omitempty"`

	// 云服务类型
	Provider *string `json:"provider,omitempty"`

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 统计体
	Counts         *[]ResourceCountItem `json:"counts,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowResourcesSummaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourcesSummaryResponse struct{}"
	}

	return strings.Join([]string{"ShowResourcesSummaryResponse", string(data)}, " ")
}
