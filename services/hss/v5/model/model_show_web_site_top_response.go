package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWebSiteTopResponse Response Object
type ShowWebSiteTopResponse struct {

	// **参数解释**: TOP5总数 **取值范围**: 取值0-5位
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: TOP5列表
	DataList       *[]CommonTopResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowWebSiteTopResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWebSiteTopResponse struct{}"
	}

	return strings.Join([]string{"ShowWebSiteTopResponse", string(data)}, " ")
}
