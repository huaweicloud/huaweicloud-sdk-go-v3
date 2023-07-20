package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportSqlJobResultResponse Response Object
type ExportSqlJobResultResponse struct {

	// 请求发送是否成功。“true”表示请求发送成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message *string `json:"message,omitempty"`

	// 此SQL将生成并提交一个新的作业，返回作业ID。用户可以使用作业ID来查询作业状态和获取作业结果。
	JobId *string `json:"job_id,omitempty"`

	// 作业执行方式，是同步还是异步执行
	JobMode        *string `json:"job_mode,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportSqlJobResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportSqlJobResultResponse struct{}"
	}

	return strings.Join([]string{"ExportSqlJobResultResponse", string(data)}, " ")
}
