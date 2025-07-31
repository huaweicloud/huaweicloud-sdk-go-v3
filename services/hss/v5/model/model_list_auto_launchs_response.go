package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAutoLaunchsResponse Response Object
type ListAutoLaunchsResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**： 自启动项列表 **取值范围**： 不涉及
	DataList       *[]AutoLauchResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListAutoLaunchsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAutoLaunchsResponse struct{}"
	}

	return strings.Join([]string{"ListAutoLaunchsResponse", string(data)}, " ")
}
