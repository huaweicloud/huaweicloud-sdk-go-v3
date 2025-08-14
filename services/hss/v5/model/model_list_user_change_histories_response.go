package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserChangeHistoriesResponse Response Object
type ListUserChangeHistoriesResponse struct {

	// **参数解释**: 账号变动总数 **取值范围**: 最小值0，最大值10000000
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 账号历史变动记录列表 **取值范围**: 最小值0，最大值200
	DataList       *[]UserChangeHistoryResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o ListUserChangeHistoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserChangeHistoriesResponse struct{}"
	}

	return strings.Join([]string{"ListUserChangeHistoriesResponse", string(data)}, " ")
}
