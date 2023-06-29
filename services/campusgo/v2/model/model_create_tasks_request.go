package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTasksRequest Request Object
type CreateTasksRequest struct {

	// 服务API,具体API名称见园区智能体提供的API参考列表中URI描述[API列表](https://support.huaweicloud.com/api-campusgo/campusgo_03_0013.html)
	ServiceName string `json:"service_name"`

	Body *CreateTasksRequestBody `json:"body,omitempty"`
}

func (o CreateTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTasksRequest struct{}"
	}

	return strings.Join([]string{"CreateTasksRequest", string(data)}, " ")
}
