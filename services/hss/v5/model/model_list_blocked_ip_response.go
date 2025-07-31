package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBlockedIpResponse Response Object
type ListBlockedIpResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// 已拦截IP详情
	DataList       *[]BlockedIpResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListBlockedIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBlockedIpResponse struct{}"
	}

	return strings.Join([]string{"ListBlockedIpResponse", string(data)}, " ")
}
