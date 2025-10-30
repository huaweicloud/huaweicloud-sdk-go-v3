package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGlobalImageAppsResponse Response Object
type ListGlobalImageAppsResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483547
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 列表 **取值范围**: 最小值0，最大值2147483647
	DataList       *[]ImageAppResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListGlobalImageAppsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalImageAppsResponse struct{}"
	}

	return strings.Join([]string{"ListGlobalImageAppsResponse", string(data)}, " ")
}
