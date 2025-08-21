package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadBlobsRawResponse Response Object
type DownloadBlobsRawResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadBlobsRawResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadBlobsRawResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadBlobsRawResponse struct{}"
	}

	return strings.Join([]string{"DownloadBlobsRawResponse", string(data)}, " ")
}
