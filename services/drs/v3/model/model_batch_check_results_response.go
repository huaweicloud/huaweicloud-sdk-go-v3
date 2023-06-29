package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCheckResultsResponse Response Object
type BatchCheckResultsResponse struct {

	// 批量查询预检查结果响应体集合。
	Results *[]QueryPreCheckResp `json:"results,omitempty"`

	// 总记录数。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchCheckResultsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCheckResultsResponse struct{}"
	}

	return strings.Join([]string{"BatchCheckResultsResponse", string(data)}, " ")
}
