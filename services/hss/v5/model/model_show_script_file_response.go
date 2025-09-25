package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// ShowScriptFileResponse Response Object
type ShowScriptFileResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o ShowScriptFileResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o ShowScriptFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScriptFileResponse struct{}"
	}

	return strings.Join([]string{"ShowScriptFileResponse", string(data)}, " ")
}
