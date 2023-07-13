package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// ShowPointTemplateResponse Response Object
type ShowPointTemplateResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o ShowPointTemplateResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o ShowPointTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPointTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowPointTemplateResponse", string(data)}, " ")
}
