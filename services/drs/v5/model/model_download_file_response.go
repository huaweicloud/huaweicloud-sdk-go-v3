package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadFileResponse Response Object
type DownloadFileResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadFileResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadFileResponse struct{}"
	}

	return strings.Join([]string{"DownloadFileResponse", string(data)}, " ")
}
