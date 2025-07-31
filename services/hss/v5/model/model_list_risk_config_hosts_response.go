package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRiskConfigHostsResponse Response Object
type ListRiskConfigHostsResponse struct {

	// **参数解释**: 受配置检测影响的服务器数据总量 **取值范围**: 不涉及
	TotalNum *int64 `json:"total_num,omitempty"`

	// **参数解释**: 数据列表 **取值范围**: 不涉及
	DataList       *[]SecurityCheckHostInfoResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o ListRiskConfigHostsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRiskConfigHostsResponse struct{}"
	}

	return strings.Join([]string{"ListRiskConfigHostsResponse", string(data)}, " ")
}
