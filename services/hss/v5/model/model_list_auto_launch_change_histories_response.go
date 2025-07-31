package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAutoLaunchChangeHistoriesResponse Response Object
type ListAutoLaunchChangeHistoriesResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// 软件历史变动记录列表
	DataList       *[]AutoLaunchChangeResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ListAutoLaunchChangeHistoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAutoLaunchChangeHistoriesResponse struct{}"
	}

	return strings.Join([]string{"ListAutoLaunchChangeHistoriesResponse", string(data)}, " ")
}
