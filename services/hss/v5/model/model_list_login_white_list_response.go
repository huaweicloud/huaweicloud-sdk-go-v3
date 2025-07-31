package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLoginWhiteListResponse Response Object
type ListLoginWhiteListResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// 登录白名单详情
	DataList       *[]LoginWhiteListResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListLoginWhiteListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLoginWhiteListResponse struct{}"
	}

	return strings.Join([]string{"ListLoginWhiteListResponse", string(data)}, " ")
}
