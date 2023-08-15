package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// ExportMqsInstanceTopicResponse Response Object
type ExportMqsInstanceTopicResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o ExportMqsInstanceTopicResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o ExportMqsInstanceTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportMqsInstanceTopicResponse struct{}"
	}

	return strings.Join([]string{"ExportMqsInstanceTopicResponse", string(data)}, " ")
}
