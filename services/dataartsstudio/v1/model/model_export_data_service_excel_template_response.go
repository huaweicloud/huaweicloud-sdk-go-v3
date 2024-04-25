package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// ExportDataServiceExcelTemplateResponse Response Object
type ExportDataServiceExcelTemplateResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o ExportDataServiceExcelTemplateResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o ExportDataServiceExcelTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportDataServiceExcelTemplateResponse struct{}"
	}

	return strings.Join([]string{"ExportDataServiceExcelTemplateResponse", string(data)}, " ")
}
