package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIgnoreHostsResponse Response Object
type ListIgnoreHostsResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释** 已忽略主机信息列表 **取值范围** 数组长度0-200
	DataList       *[]IgnoreHostResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListIgnoreHostsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIgnoreHostsResponse struct{}"
	}

	return strings.Join([]string{"ListIgnoreHostsResponse", string(data)}, " ")
}
