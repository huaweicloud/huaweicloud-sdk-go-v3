package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQualityTaskResponse Response Object
type ListQualityTaskResponse struct {

	// 总条数
	Count *int64 `json:"count,omitempty"`

	// 分页数据
	Resources      *[]QualityTaskOverviewVo `json:"resources,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListQualityTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQualityTaskResponse struct{}"
	}

	return strings.Join([]string{"ListQualityTaskResponse", string(data)}, " ")
}
