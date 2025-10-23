package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReportResourceDetailsDataResponse Response Object
type ShowReportResourceDetailsDataResponse struct {

	// 本次分页查询响应的数据条数
	Count *int64 `json:"count,omitempty"`

	// 下一次分页查询的游标
	NextMarker *string `json:"next_marker,omitempty"`

	// 资源列表
	Resources      *[]ResourceEntity `json:"resources,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowReportResourceDetailsDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReportResourceDetailsDataResponse struct{}"
	}

	return strings.Join([]string{"ShowReportResourceDetailsDataResponse", string(data)}, " ")
}
