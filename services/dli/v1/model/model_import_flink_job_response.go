package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportFlinkJobResponse Response Object
type ImportFlinkJobResponse struct {

	// 执行请求是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 消息内容。
	Message *string `json:"message,omitempty"`

	// 作业导入结果。
	JobMapping     *[]JobMapInfo `json:"job_mapping,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ImportFlinkJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportFlinkJobResponse struct{}"
	}

	return strings.Join([]string{"ImportFlinkJobResponse", string(data)}, " ")
}
