package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWeakPasswordUsersResponse Response Object
type ListWeakPasswordUsersResponse struct {

	// **参数解释**: 弱口令总数 **取值范围**: 不涉及
	TotalNum *int64 `json:"total_num,omitempty"`

	// **参数解释**: 弱口令列表 **取值范围**: 不涉及
	DataList       *[]WeakPwdListInfoResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListWeakPasswordUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWeakPasswordUsersResponse struct{}"
	}

	return strings.Join([]string{"ListWeakPasswordUsersResponse", string(data)}, " ")
}
