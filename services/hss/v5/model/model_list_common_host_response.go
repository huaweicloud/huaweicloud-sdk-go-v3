package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCommonHostResponse Response Object
type ListCommonHostResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// 列表
	DataList       *[]CommonHostResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListCommonHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCommonHostResponse struct{}"
	}

	return strings.Join([]string{"ListCommonHostResponse", string(data)}, " ")
}
