package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadTemplateResponse Response Object
type DownloadTemplateResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadTemplateResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadTemplateResponse struct{}"
	}

	return strings.Join([]string{"DownloadTemplateResponse", string(data)}, " ")
}
