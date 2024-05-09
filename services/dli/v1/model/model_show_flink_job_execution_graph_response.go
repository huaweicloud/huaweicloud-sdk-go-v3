package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlinkJobExecutionGraphResponse Response Object
type ShowFlinkJobExecutionGraphResponse struct {

	// 请求是否成功
	IsSuccess *string `json:"is_success,omitempty"`

	// 消息内容。
	Message *string `json:"message,omitempty"`

	ExecuteGraph   *FlinkJobExecutionGraph `json:"execute_graph,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ShowFlinkJobExecutionGraphResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlinkJobExecutionGraphResponse struct{}"
	}

	return strings.Join([]string{"ShowFlinkJobExecutionGraphResponse", string(data)}, " ")
}
