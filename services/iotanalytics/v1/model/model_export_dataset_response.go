package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// Response Object
type ExportDatasetResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o ExportDatasetResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o ExportDatasetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportDatasetResponse struct{}"
	}

	return strings.Join([]string{"ExportDatasetResponse", string(data)}, " ")
}
