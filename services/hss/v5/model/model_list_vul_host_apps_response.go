package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVulHostAppsResponse Response Object
type ListVulHostAppsResponse struct {

	// **参数解释**: 软件信息列表 **取值范围**: 不涉及
	DataList *[]VulHostAppsResponseInfoDataList `json:"data_list,omitempty"`

	// **参数解释**: 总数 **取值范围**: 取值0-2147483647
	TotalNum       *int32 `json:"total_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListVulHostAppsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulHostAppsResponse struct{}"
	}

	return strings.Join([]string{"ListVulHostAppsResponse", string(data)}, " ")
}
