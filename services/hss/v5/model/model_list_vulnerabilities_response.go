package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVulnerabilitiesResponse Response Object
type ListVulnerabilitiesResponse struct {

	// **参数解释**: 漏洞总数 **取值范围**: 取值0-2147483647
	TotalNum *int64 `json:"total_num,omitempty"`

	// **参数解释**: 漏洞数据列表 **取值范围**: 不涉及
	DataList       *[]VulInfo `json:"data_list,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListVulnerabilitiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulnerabilitiesResponse struct{}"
	}

	return strings.Join([]string{"ListVulnerabilitiesResponse", string(data)}, " ")
}
