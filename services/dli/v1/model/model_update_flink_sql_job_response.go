package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFlinkSqlJobResponse Response Object
type UpdateFlinkSqlJobResponse struct {

	// 执行请求是否成功。“true”表示请求执行成功。
	IsSuccess *string `json:"is_success,omitempty"`

	// 消息内容。
	Message *string `json:"message,omitempty"`

	Job            *FlinkJobUpdateTime `json:"job,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o UpdateFlinkSqlJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFlinkSqlJobResponse struct{}"
	}

	return strings.Join([]string{"UpdateFlinkSqlJobResponse", string(data)}, " ")
}
