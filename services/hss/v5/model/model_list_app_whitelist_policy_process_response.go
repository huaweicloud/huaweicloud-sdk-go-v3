package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppWhitelistPolicyProcessResponse Response Object
type ListAppWhitelistPolicyProcessResponse struct {

	// data list
	DataList *[]AppWhitelistPolicyProcessResponseInfo `json:"data_list,omitempty"`

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum       *int32 `json:"total_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAppWhitelistPolicyProcessResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppWhitelistPolicyProcessResponse struct{}"
	}

	return strings.Join([]string{"ListAppWhitelistPolicyProcessResponse", string(data)}, " ")
}
