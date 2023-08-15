package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadFailureReportResponse Response Object
type DownloadFailureReportResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadFailureReportResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadFailureReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadFailureReportResponse struct{}"
	}

	return strings.Join([]string{"DownloadFailureReportResponse", string(data)}, " ")
}
