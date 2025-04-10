package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRaspPoliciesResponse Response Object
type ListRaspPoliciesResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 策略信息列表
	DataList       *[]PolicyInfo `json:"data_list,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListRaspPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRaspPoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListRaspPoliciesResponse", string(data)}, " ")
}
