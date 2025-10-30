package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVulRepairCmdsResponse Response Object
type ListVulRepairCmdsResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值10000
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 漏洞修复命令列表 **取值范围**: 最少0条，最多2147483647条
	DataList       *[]VulRepairCmdInfo `json:"data_list,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListVulRepairCmdsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulRepairCmdsResponse struct{}"
	}

	return strings.Join([]string{"ListVulRepairCmdsResponse", string(data)}, " ")
}
