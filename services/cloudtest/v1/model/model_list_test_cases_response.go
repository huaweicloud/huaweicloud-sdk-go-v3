package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTestCasesResponse Response Object
type ListTestCasesResponse struct {

	// 用例详情列表
	Values *[]ExternalTestCaseVo `json:"values,omitempty"`

	// 用例总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTestCasesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTestCasesResponse struct{}"
	}

	return strings.Join([]string{"ListTestCasesResponse", string(data)}, " ")
}
