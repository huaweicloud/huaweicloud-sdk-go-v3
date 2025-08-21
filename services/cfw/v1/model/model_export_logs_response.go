package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// ExportLogsResponse Response Object
type ExportLogsResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o ExportLogsResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o ExportLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportLogsResponse struct{}"
	}

	return strings.Join([]string{"ExportLogsResponse", string(data)}, " ")
}
