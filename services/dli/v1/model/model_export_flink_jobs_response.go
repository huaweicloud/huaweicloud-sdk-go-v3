package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportFlinkJobsResponse Response Object
type ExportFlinkJobsResponse struct {

	// 执行请求是否成功。“true”表示请求执行成功。
	IsSuccess *string `json:"is_success,omitempty"`

	// 消息内容
	Message *string `json:"message,omitempty"`

	// OBS上导出作业zip文件名。
	ZipFile        *[]string `json:"zip_file,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ExportFlinkJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportFlinkJobsResponse struct{}"
	}

	return strings.Join([]string{"ExportFlinkJobsResponse", string(data)}, " ")
}
