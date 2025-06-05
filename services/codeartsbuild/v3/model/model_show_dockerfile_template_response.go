package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// ShowDockerfileTemplateResponse Response Object
type ShowDockerfileTemplateResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o ShowDockerfileTemplateResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o ShowDockerfileTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDockerfileTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowDockerfileTemplateResponse", string(data)}, " ")
}
