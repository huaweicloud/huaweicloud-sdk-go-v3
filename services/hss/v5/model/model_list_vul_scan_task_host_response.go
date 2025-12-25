package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVulScanTaskHostResponse Response Object
type ListVulScanTaskHostResponse struct {

	// **参数解释**: 总数 **取值范围**: 取值0-2147483647
	TotalNum *int64 `json:"total_num,omitempty"`

	// **参数解释**: 漏洞扫描任务对应的主机列表 **取值范围**: 最小值0，最大值2147483647
	DataList       *[]VulScanTaskHostInfo `json:"data_list,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListVulScanTaskHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulScanTaskHostResponse struct{}"
	}

	return strings.Join([]string{"ListVulScanTaskHostResponse", string(data)}, " ")
}
