package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// Response Object
type DownloadDataTraceResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadDataTraceResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadDataTraceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadDataTraceResponse struct{}"
	}

	return strings.Join([]string{"DownloadDataTraceResponse", string(data)}, " ")
}
