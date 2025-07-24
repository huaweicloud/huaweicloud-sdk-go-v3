package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadCreateTemplateResponse Response Object
type DownloadCreateTemplateResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadCreateTemplateResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadCreateTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadCreateTemplateResponse struct{}"
	}

	return strings.Join([]string{"DownloadCreateTemplateResponse", string(data)}, " ")
}
