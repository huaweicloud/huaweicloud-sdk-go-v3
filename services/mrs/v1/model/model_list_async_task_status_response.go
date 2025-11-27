package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAsyncTaskStatusResponse Response Object
type ListAsyncTaskStatusResponse struct {

	// 获取集群委托切换任务状态 - running：集群委托切换中 - success：集群委托切换成功
	Status         *string `json:"Status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListAsyncTaskStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAsyncTaskStatusResponse struct{}"
	}

	return strings.Join([]string{"ListAsyncTaskStatusResponse", string(data)}, " ")
}
