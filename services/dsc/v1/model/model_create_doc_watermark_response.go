package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// Response Object
type CreateDocWatermarkResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o CreateDocWatermarkResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o CreateDocWatermarkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDocWatermarkResponse struct{}"
	}

	return strings.Join([]string{"CreateDocWatermarkResponse", string(data)}, " ")
}
