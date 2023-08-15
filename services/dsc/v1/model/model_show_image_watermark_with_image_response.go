package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// ShowImageWatermarkWithImageResponse Response Object
type ShowImageWatermarkWithImageResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o ShowImageWatermarkWithImageResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o ShowImageWatermarkWithImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageWatermarkWithImageResponse struct{}"
	}

	return strings.Join([]string{"ShowImageWatermarkWithImageResponse", string(data)}, " ")
}
