package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// ExportUserLoginInfoNewResponse Response Object
type ExportUserLoginInfoNewResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o ExportUserLoginInfoNewResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o ExportUserLoginInfoNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportUserLoginInfoNewResponse struct{}"
	}

	return strings.Join([]string{"ExportUserLoginInfoNewResponse", string(data)}, " ")
}
