package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTaskExecuteDetailResponse Response Object
type ListTaskExecuteDetailResponse struct {

	// 总数。
	Count *int32 `json:"count,omitempty"`

	// 服务器组列表，返回列表条目数量上限为分页的最大上限值。
	Items          *[]ScheduleTaskExecuteDetail `json:"items,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListTaskExecuteDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskExecuteDetailResponse struct{}"
	}

	return strings.Join([]string{"ListTaskExecuteDetailResponse", string(data)}, " ")
}
