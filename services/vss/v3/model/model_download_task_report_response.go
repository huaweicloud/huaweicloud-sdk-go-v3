package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadTaskReportResponse Response Object
type DownloadTaskReportResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadTaskReportResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadTaskReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadTaskReportResponse struct{}"
	}

	return strings.Join([]string{"DownloadTaskReportResponse", string(data)}, " ")
}
