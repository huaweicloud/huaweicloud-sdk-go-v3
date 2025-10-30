package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEventTypeResponse Response Object
type ListEventTypeResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// 各种类型详情
	DataList       *[]EventTypeDetailResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListEventTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventTypeResponse struct{}"
	}

	return strings.Join([]string{"ListEventTypeResponse", string(data)}, " ")
}
