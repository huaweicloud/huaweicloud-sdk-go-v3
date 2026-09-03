package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadHttpCertResponse Response Object
type DownloadHttpCertResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadHttpCertResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadHttpCertResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadHttpCertResponse struct{}"
	}

	return strings.Join([]string{"DownloadHttpCertResponse", string(data)}, " ")
}
