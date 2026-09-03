package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadHttpSignCertResponse Response Object
type DownloadHttpSignCertResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadHttpSignCertResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadHttpSignCertResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadHttpSignCertResponse struct{}"
	}

	return strings.Join([]string{"DownloadHttpSignCertResponse", string(data)}, " ")
}
