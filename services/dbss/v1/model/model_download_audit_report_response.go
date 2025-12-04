package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadAuditReportResponse Response Object
type DownloadAuditReportResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadAuditReportResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadAuditReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadAuditReportResponse struct{}"
	}

	return strings.Join([]string{"DownloadAuditReportResponse", string(data)}, " ")
}
