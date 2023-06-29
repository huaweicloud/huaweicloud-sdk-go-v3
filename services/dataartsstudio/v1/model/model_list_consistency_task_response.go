package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConsistencyTaskResponse Response Object
type ListConsistencyTaskResponse struct {

	// 总条数
	Count *int64 `json:"count,omitempty"`

	// 分页数据
	Resources      *[]QualityTaskOverviewVo2 `json:"resources,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListConsistencyTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConsistencyTaskResponse struct{}"
	}

	return strings.Join([]string{"ListConsistencyTaskResponse", string(data)}, " ")
}
