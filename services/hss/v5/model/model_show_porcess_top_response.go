package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPorcessTopResponse Response Object
type ShowPorcessTopResponse struct {

	// **参数解释**: TOP5总数 **取值范围**: 取值0-5位
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: TOP5列表
	DataList       *[]CommonTopResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowPorcessTopResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPorcessTopResponse struct{}"
	}

	return strings.Join([]string{"ShowPorcessTopResponse", string(data)}, " ")
}
