package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadArchiveResponse Response Object
type DownloadArchiveResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadArchiveResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadArchiveResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadArchiveResponse struct{}"
	}

	return strings.Join([]string{"DownloadArchiveResponse", string(data)}, " ")
}
