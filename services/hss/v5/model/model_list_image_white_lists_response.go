package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageWhiteListsResponse Response Object
type ListImageWhiteListsResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 白名单列表 **取值范围**: 最小值0，最大值200
	DataList       *[]ImageWhiteListDetailResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ListImageWhiteListsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageWhiteListsResponse struct{}"
	}

	return strings.Join([]string{"ListImageWhiteListsResponse", string(data)}, " ")
}
