package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchListRposAndRtosResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty" xml:"count"`

	// 批量查询RPO&RTO的响应体集合
	Results        *[]QueryRpoAndRtoResp `json:"results,omitempty" xml:"results"`
	HttpStatusCode int                   `json:"-"`
}

func (o BatchListRposAndRtosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListRposAndRtosResponse struct{}"
	}

	return strings.Join([]string{"BatchListRposAndRtosResponse", string(data)}, " ")
}
