package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainClientStatsResponse Response Object
type ListDomainClientStatsResponse struct {

	// 区域
	ServiceArea *string `json:"service_area,omitempty"`

	// 按域名维每天客户端访问详情统计
	Result         *[]map[string]interface{} `json:"result,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListDomainClientStatsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainClientStatsResponse struct{}"
	}

	return strings.Join([]string{"ListDomainClientStatsResponse", string(data)}, " ")
}
