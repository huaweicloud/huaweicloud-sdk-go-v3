package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPasswordComplexityResponse Response Object
type ListPasswordComplexityResponse struct {

	// **参数解释**: 记录总数 **取值范围**: 不涉及
	TotalNum *int64 `json:"total_num,omitempty"`

	// **参数解释**: 口令复杂度策略检测列表 **取值范围**: 不涉及
	DataList       *[]PwdPolicyInfoResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListPasswordComplexityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPasswordComplexityResponse struct{}"
	}

	return strings.Join([]string{"ListPasswordComplexityResponse", string(data)}, " ")
}
