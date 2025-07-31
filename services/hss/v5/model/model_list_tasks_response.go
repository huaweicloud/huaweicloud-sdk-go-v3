package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTasksResponse Response Object
type ListTasksResponse struct {

	// 任务列表数量
	TotalNum *int32 `json:"total_num,omitempty"`

	// 任务列表
	DataList       *[]ListTasksResponseInfoDataList `json:"data_list,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o ListTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTasksResponse struct{}"
	}

	return strings.Join([]string{"ListTasksResponse", string(data)}, " ")
}
