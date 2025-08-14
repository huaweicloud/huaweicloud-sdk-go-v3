package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadResultFileResponse Response Object
type DownloadResultFileResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadResultFileResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadResultFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadResultFileResponse struct{}"
	}

	return strings.Join([]string{"DownloadResultFileResponse", string(data)}, " ")
}
