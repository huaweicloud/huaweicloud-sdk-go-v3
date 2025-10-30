package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVulWhiteListResponse Response Object
type ListVulWhiteListResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 漏洞白名单列表 **取值范围**: 最小值0，最大值2147483647
	DataList       *[]VulWhiteListDetailResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ListVulWhiteListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulWhiteListResponse struct{}"
	}

	return strings.Join([]string{"ListVulWhiteListResponse", string(data)}, " ")
}
