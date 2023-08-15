package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadAppVersionResponse Response Object
type DownloadAppVersionResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadAppVersionResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadAppVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadAppVersionResponse struct{}"
	}

	return strings.Join([]string{"DownloadAppVersionResponse", string(data)}, " ")
}
