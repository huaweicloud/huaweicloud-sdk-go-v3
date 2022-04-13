package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// Response Object
type DownloadApplicationCodeResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadApplicationCodeResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadApplicationCodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadApplicationCodeResponse struct{}"
	}

	return strings.Join([]string{"DownloadApplicationCodeResponse", string(data)}, " ")
}
