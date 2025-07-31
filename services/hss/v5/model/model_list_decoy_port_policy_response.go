package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDecoyPortPolicyResponse Response Object
type ListDecoyPortPolicyResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// policy list
	DataList       *[]PolicyListDataList `json:"data_list,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListDecoyPortPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDecoyPortPolicyResponse struct{}"
	}

	return strings.Join([]string{"ListDecoyPortPolicyResponse", string(data)}, " ")
}
