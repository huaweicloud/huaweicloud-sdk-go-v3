package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDatabaseResponse Response Object
type DeleteDatabaseResponse struct {

	// 执行请求是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息为空。
	Message *string `json:"message,omitempty"`

	// 异步删除数据库的时候返回的删除作业ID
	JobId *string `json:"job_id,omitempty"`

	// 删除数据库的时候的方式，是同步删除还是异步删除
	JobMode *string `json:"job_mode,omitempty"`

	// 异步执行作业时返回执行删除作业类型
	JobType        *string `json:"job_type,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDatabaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatabaseResponse struct{}"
	}

	return strings.Join([]string{"DeleteDatabaseResponse", string(data)}, " ")
}
