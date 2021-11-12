package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SlowlogDownloadRequest struct {
	// - 请求ID，uuid，代表此次获取慢日志的请求ID。

	RequestId string `json:"request_id"`
	// - 需要下载的文件的文件名, 当引擎为SQL Server时为必选。

	FileName *string `json:"file_name,omitempty"`
}

func (o SlowlogDownloadRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowlogDownloadRequest struct{}"
	}

	return strings.Join([]string{"SlowlogDownloadRequest", string(data)}, " ")
}
