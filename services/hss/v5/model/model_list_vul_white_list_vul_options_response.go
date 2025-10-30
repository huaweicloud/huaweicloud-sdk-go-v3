package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVulWhiteListVulOptionsResponse Response Object
type ListVulWhiteListVulOptionsResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 漏洞列表 **取值范围**: 最小值0，最大值2147483647
	DataList       *[]VulWhiteListVulOptionsResponseInfoDataList `json:"data_list,omitempty"`
	HttpStatusCode int                                           `json:"-"`
}

func (o ListVulWhiteListVulOptionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulWhiteListVulOptionsResponse struct{}"
	}

	return strings.Join([]string{"ListVulWhiteListVulOptionsResponse", string(data)}, " ")
}
