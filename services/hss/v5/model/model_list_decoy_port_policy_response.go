package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDecoyPortPolicyResponse Response Object
type ListDecoyPortPolicyResponse struct {

	// 总数
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
