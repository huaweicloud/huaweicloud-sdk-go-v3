package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFlinkSqlJobResponse Response Object
type CreateFlinkSqlJobResponse struct {

	// 执行请求是否成功。“true”表示请求执行成功。
	IsSuccess *string `json:"is_success,omitempty"`

	// 消息内容
	Message *string `json:"message,omitempty"`

	Job            *FlinkJobStatus `json:"job,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o CreateFlinkSqlJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlinkSqlJobResponse struct{}"
	}

	return strings.Join([]string{"CreateFlinkSqlJobResponse", string(data)}, " ")
}
