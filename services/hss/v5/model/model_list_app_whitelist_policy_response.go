package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppWhitelistPolicyResponse Response Object
type ListAppWhitelistPolicyResponse struct {

	// data list
	DataList *[]AppWhitelistPolicyResponseInfo `json:"data_list,omitempty"`

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum       *int32 `json:"total_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAppWhitelistPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppWhitelistPolicyResponse struct{}"
	}

	return strings.Join([]string{"ListAppWhitelistPolicyResponse", string(data)}, " ")
}
