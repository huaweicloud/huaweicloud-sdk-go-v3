package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadHotKeyResponse Response Object
type DownloadHotKeyResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadHotKeyResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadHotKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadHotKeyResponse struct{}"
	}

	return strings.Join([]string{"DownloadHotKeyResponse", string(data)}, " ")
}
