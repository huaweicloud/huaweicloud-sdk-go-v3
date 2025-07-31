package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWebTamperHostResponse Response Object
type ListWebTamperHostResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// data list
	DataList       *[]WebTamperHostResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListWebTamperHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWebTamperHostResponse struct{}"
	}

	return strings.Join([]string{"ListWebTamperHostResponse", string(data)}, " ")
}
