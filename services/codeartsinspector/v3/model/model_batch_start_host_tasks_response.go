package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchStartHostTasksResponse Response Object
type BatchStartHostTasksResponse struct {

	// 扫描任务ID对应主机id列表
	Results        *[]BatchStartHostTasksResponseBodyResults `json:"results,omitempty"`
	HttpStatusCode int                                       `json:"-"`
}

func (o BatchStartHostTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartHostTasksResponse struct{}"
	}

	return strings.Join([]string{"BatchStartHostTasksResponse", string(data)}, " ")
}
