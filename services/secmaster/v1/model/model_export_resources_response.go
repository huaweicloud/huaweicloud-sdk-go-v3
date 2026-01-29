package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// ExportResourcesResponse Response Object
type ExportResourcesResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o ExportResourcesResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o ExportResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportResourcesResponse struct{}"
	}

	return strings.Join([]string{"ExportResourcesResponse", string(data)}, " ")
}
