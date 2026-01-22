package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadImportTemplateResponse Response Object
type DownloadImportTemplateResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadImportTemplateResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadImportTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadImportTemplateResponse struct{}"
	}

	return strings.Join([]string{"DownloadImportTemplateResponse", string(data)}, " ")
}
