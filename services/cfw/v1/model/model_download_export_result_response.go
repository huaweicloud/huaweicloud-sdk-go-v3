package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadExportResultResponse Response Object
type DownloadExportResultResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadExportResultResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadExportResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadExportResultResponse struct{}"
	}

	return strings.Join([]string{"DownloadExportResultResponse", string(data)}, " ")
}
