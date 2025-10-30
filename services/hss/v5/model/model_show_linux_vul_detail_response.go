package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLinuxVulDetailResponse Response Object
type ShowLinuxVulDetailResponse struct {

	// **参数解释**: 记录总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: linux漏洞cve列表 **取值范围**: 最小值0，最大值10000
	DataList       *[]LinuxVulDetailInfo `json:"data_list,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowLinuxVulDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLinuxVulDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowLinuxVulDetailResponse", string(data)}, " ")
}
