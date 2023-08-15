package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlinkExecuteGraphResponse Response Object
type ShowFlinkExecuteGraphResponse struct {

	// 请求是否成功
	IsSuccess *bool `json:"is_success,omitempty"`

	// 消息内容。
	Message *string `json:"message,omitempty"`

	ExecuteGraph   *StreamGraphInfo `json:"execute_graph,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowFlinkExecuteGraphResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlinkExecuteGraphResponse struct{}"
	}

	return strings.Join([]string{"ShowFlinkExecuteGraphResponse", string(data)}, " ")
}
