package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHostRaspProtectHistoryInfoResponse Response Object
type ListHostRaspProtectHistoryInfoResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 动态网页防篡改防护事件列表 **取值范围**: 最小值0，最大值200
	DataList       *[]HostRaspProtectHistoryResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o ListHostRaspProtectHistoryInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostRaspProtectHistoryInfoResponse struct{}"
	}

	return strings.Join([]string{"ListHostRaspProtectHistoryInfoResponse", string(data)}, " ")
}
