package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWebSiteInfoResponse Response Object
type ListWebSiteInfoResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: Web站点列表 **取值范围**: 最小值0，最大值10000
	DataList       *[]WebSiteInfo `json:"data_list,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListWebSiteInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWebSiteInfoResponse struct{}"
	}

	return strings.Join([]string{"ListWebSiteInfoResponse", string(data)}, " ")
}
