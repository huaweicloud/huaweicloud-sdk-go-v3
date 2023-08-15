package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTestCaseHistoriesResponse Response Object
type ListTestCaseHistoriesResponse struct {

	// 起始记录数 大于 实际总条数时， 值为0， 分页请求才有此值
	Total *int32 `json:"total,omitempty"`

	// 实际的数据类型：单个对象，集合 或 NULL
	Values         *[]ExternalTestCaseHistoryVo `json:"values,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListTestCaseHistoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTestCaseHistoriesResponse struct{}"
	}

	return strings.Join([]string{"ListTestCaseHistoriesResponse", string(data)}, " ")
}
