package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// ExportAopworkflowResponse Response Object
type ExportAopworkflowResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o ExportAopworkflowResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o ExportAopworkflowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportAopworkflowResponse struct{}"
	}

	return strings.Join([]string{"ExportAopworkflowResponse", string(data)}, " ")
}
