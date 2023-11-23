package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFunctionApplicationsResponse Response Object
type ListFunctionApplicationsResponse struct {

	// 函数应用列表
	Applications *[]ListFunctionApplicationResult `json:"applications,omitempty"`

	// 下次读取位置
	NextMarker *int64 `json:"next_marker,omitempty"`

	// 应用程序总数
	Count          *int64 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListFunctionApplicationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFunctionApplicationsResponse struct{}"
	}

	return strings.Join([]string{"ListFunctionApplicationsResponse", string(data)}, " ")
}
