package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCicdConfigurationsResponse Response Object
type ListCicdConfigurationsResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// CI/CD配置列表
	DataList       *[]CicdConfigurationsResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ListCicdConfigurationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCicdConfigurationsResponse struct{}"
	}

	return strings.Join([]string{"ListCicdConfigurationsResponse", string(data)}, " ")
}
