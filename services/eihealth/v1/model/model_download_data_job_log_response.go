package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// Response Object
type DownloadDataJobLogResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadDataJobLogResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadDataJobLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadDataJobLogResponse struct{}"
	}

	return strings.Join([]string{"DownloadDataJobLogResponse", string(data)}, " ")
}
