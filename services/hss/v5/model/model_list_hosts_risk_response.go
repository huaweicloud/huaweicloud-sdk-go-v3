package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHostsRiskResponse Response Object
type ListHostsRiskResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 查询ECS风险列表
	DataList       *[]HostRiskResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListHostsRiskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostsRiskResponse struct{}"
	}

	return strings.Join([]string{"ListHostsRiskResponse", string(data)}, " ")
}
