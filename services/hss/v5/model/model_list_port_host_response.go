package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPortHostResponse Response Object
type ListPortHostResponse struct {

	// **参数解释**: 机器总数 **取值范围**: 最小值0， 最大值10000
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 机器信息列表 **取值范围**: 最小值0， 最大值10000
	DataList       *[]PortHostResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListPortHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPortHostResponse struct{}"
	}

	return strings.Join([]string{"ListPortHostResponse", string(data)}, " ")
}
