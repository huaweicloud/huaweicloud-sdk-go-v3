package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFileEventsResponse Response Object
type ListFileEventsResponse struct {

	// 变更文件数量
	TotalNum *int32 `json:"total_num,omitempty"`

	// 变更文件信息列表
	DataList       *[]FileEventResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListFileEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFileEventsResponse struct{}"
	}

	return strings.Join([]string{"ListFileEventsResponse", string(data)}, " ")
}
