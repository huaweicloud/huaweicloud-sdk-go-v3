package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// ShowTemplateResponse Response Object
type ShowTemplateResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o ShowTemplateResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o ShowTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowTemplateResponse", string(data)}, " ")
}
