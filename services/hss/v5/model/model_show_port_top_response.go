package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPortTopResponse Response Object
type ShowPortTopResponse struct {

	// **参数解释**: TOP5总数 **取值范围**: 取值0-5位
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: TOP5列表
	DataList       *[]CommonTopResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowPortTopResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPortTopResponse struct{}"
	}

	return strings.Join([]string{"ShowPortTopResponse", string(data)}, " ")
}
