package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadResourceTemplateResponse Response Object
type DownloadResourceTemplateResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadResourceTemplateResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadResourceTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadResourceTemplateResponse struct{}"
	}

	return strings.Join([]string{"DownloadResourceTemplateResponse", string(data)}, " ")
}
