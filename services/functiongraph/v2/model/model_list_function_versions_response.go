package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListFunctionVersionsResponse struct {

	// 版本列表
	Versions *[]ListFunctionVersionResult `json:"versions,omitempty" xml:"versions"`

	// 下一次记录位置
	NextMarker *int64 `json:"next_marker,omitempty" xml:"next_marker"`

	// 版本总数
	Count          *int64 `json:"count,omitempty" xml:"count"`
	HttpStatusCode int    `json:"-"`
}

func (o ListFunctionVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFunctionVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListFunctionVersionsResponse", string(data)}, " ")
}
