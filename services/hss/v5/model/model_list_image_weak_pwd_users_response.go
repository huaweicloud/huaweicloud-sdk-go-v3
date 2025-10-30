package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageWeakPwdUsersResponse Response Object
type ListImageWeakPwdUsersResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 弱口令列表 **取值范围**: 最小值0，最大值2147483647
	DataList       *[]ImageWeakPwdListInfoResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ListImageWeakPwdUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageWeakPwdUsersResponse struct{}"
	}

	return strings.Join([]string{"ListImageWeakPwdUsersResponse", string(data)}, " ")
}
