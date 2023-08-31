package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIegResponse Response Object
type ListIegResponse struct {

	// 智能企业网关列表
	IntelligentEnterpriseGateways *[]IegItem `json:"intelligent_enterprise_gateways,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 智能企业网关数量
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListIegResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIegResponse struct{}"
	}

	return strings.Join([]string{"ListIegResponse", string(data)}, " ")
}
