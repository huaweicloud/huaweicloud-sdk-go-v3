package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchListProgressesResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty" xml:"count"`

	// 批量查询迁移进度的响应体集合
	Results        *[]QueryProgressResp `json:"results,omitempty" xml:"results"`
	HttpStatusCode int                  `json:"-"`
}

func (o BatchListProgressesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListProgressesResponse struct{}"
	}

	return strings.Join([]string{"BatchListProgressesResponse", string(data)}, " ")
}
