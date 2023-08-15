package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// CreateImageWatermarkResponse Response Object
type CreateImageWatermarkResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o CreateImageWatermarkResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o CreateImageWatermarkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageWatermarkResponse struct{}"
	}

	return strings.Join([]string{"CreateImageWatermarkResponse", string(data)}, " ")
}
