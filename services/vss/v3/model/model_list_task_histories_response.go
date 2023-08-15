package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTaskHistoriesResponse Response Object
type ListTaskHistoriesResponse struct {

	// 网站历史扫描任务总数
	Total *int32 `json:"total,omitempty"`

	// 网站历史扫描任务列表
	Data           *[]ShowTasksResponseBody `json:"data,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListTaskHistoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskHistoriesResponse struct{}"
	}

	return strings.Join([]string{"ListTaskHistoriesResponse", string(data)}, " ")
}
