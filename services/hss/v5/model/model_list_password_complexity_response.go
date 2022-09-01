package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPasswordComplexityResponse struct {

	// 记录总数
	TotalNum *int64 `json:"total_num,omitempty" xml:"total_num"`

	// 口令复杂度策略检测列表
	DataList       *[]PwdPolicyInfoResponseInfo `json:"data_list,omitempty" xml:"data_list"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListPasswordComplexityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPasswordComplexityResponse struct{}"
	}

	return strings.Join([]string{"ListPasswordComplexityResponse", string(data)}, " ")
}
