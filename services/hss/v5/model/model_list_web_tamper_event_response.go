package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWebTamperEventResponse Response Object
type ListWebTamperEventResponse struct {

	// **参数解释**: 网页防篡改事件总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 网页防篡改事件列表 **取值范围**: 取值0-200个WebTamperEventResponseInfo对象
	DataList       *[]WebTamperEventResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListWebTamperEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWebTamperEventResponse struct{}"
	}

	return strings.Join([]string{"ListWebTamperEventResponse", string(data)}, " ")
}
