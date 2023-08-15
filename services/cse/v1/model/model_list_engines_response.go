package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnginesResponse Response Object
type ListEnginesResponse struct {

	// 当前符合查询条件的微服务引擎专享版总个数
	Total *int32 `json:"total,omitempty"`

	// 微服务引擎专享版详情
	Data           *[]EngineSimpleInfo `json:"data,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListEnginesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnginesResponse struct{}"
	}

	return strings.Join([]string{"ListEnginesResponse", string(data)}, " ")
}
