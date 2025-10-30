package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLocalImageAppInfoResponse Response Object
type ListLocalImageAppInfoResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483547
	TotalNum *int32 `json:"total_num,omitempty"`

	// 列表
	DataList       *[]LocalImageAppList `json:"data_list,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListLocalImageAppInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLocalImageAppInfoResponse struct{}"
	}

	return strings.Join([]string{"ListLocalImageAppInfoResponse", string(data)}, " ")
}
