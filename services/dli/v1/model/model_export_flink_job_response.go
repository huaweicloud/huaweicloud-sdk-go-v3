package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportFlinkJobResponse Response Object
type ExportFlinkJobResponse struct {

	// 执行请求是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 消息内容
	Message *string `json:"message,omitempty"`

	// OBS上导出作业zip文件名。
	ZipFile        *[]string `json:"zip_file,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ExportFlinkJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportFlinkJobResponse struct{}"
	}

	return strings.Join([]string{"ExportFlinkJobResponse", string(data)}, " ")
}
