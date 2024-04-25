package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// ExportDataServiceExcelResponse Response Object
type ExportDataServiceExcelResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o ExportDataServiceExcelResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o ExportDataServiceExcelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportDataServiceExcelResponse struct{}"
	}

	return strings.Join([]string{"ExportDataServiceExcelResponse", string(data)}, " ")
}
