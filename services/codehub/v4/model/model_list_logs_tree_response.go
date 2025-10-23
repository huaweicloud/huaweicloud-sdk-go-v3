package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogsTreeResponse Response Object
type ListLogsTreeResponse struct {

	// 文件树列表
	Body           *[]LogTreeDto `json:"body,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListLogsTreeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogsTreeResponse struct{}"
	}

	return strings.Join([]string{"ListLogsTreeResponse", string(data)}, " ")
}
