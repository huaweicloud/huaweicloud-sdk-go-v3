package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWindowsVulDetailResponse Response Object
type ShowWindowsVulDetailResponse struct {

	// **参数解释**: 数据总条数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 软件漏洞cve列表 **取值范围**: 不涉及
	DataList       *[]WindowsVulDetailInfo `json:"data_list,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ShowWindowsVulDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWindowsVulDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowWindowsVulDetailResponse", string(data)}, " ")
}
