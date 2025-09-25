package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageAppsResponse Response Object
type ListImageAppsResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// 列表
	DataList       *[]ImageAppsInfo `json:"data_list,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListImageAppsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageAppsResponse struct{}"
	}

	return strings.Join([]string{"ListImageAppsResponse", string(data)}, " ")
}
