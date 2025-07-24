package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadServerLogsResponse Response Object
type DownloadServerLogsResponse struct {
	ContentDisposition *string `json:"Content-Disposition,omitempty"`

	ContentTransferEncoding *string `json:"Content-Transfer-Encoding,omitempty"`

	ContentType    *string       `json:"Content-Type,omitempty"`
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadServerLogsResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadServerLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadServerLogsResponse struct{}"
	}

	return strings.Join([]string{"DownloadServerLogsResponse", string(data)}, " ")
}
