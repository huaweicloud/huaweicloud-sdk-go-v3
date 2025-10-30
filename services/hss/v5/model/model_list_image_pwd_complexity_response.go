package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImagePwdComplexityResponse Response Object
type ListImagePwdComplexityResponse struct {

	// **参数解释**: 记录总数 **取值范围**: 最小值0，2147483647
	TotalNum *int64 `json:"total_num,omitempty"`

	// **参数解释**: 口令复杂度策略检测列表 **取值范围**: 最小值0，最大值2147483647
	DataList       *[]ImagePwdComplexityInfoResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o ListImagePwdComplexityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImagePwdComplexityResponse struct{}"
	}

	return strings.Join([]string{"ListImagePwdComplexityResponse", string(data)}, " ")
}
