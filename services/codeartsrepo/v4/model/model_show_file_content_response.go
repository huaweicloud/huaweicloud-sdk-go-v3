package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// ShowFileContentResponse Response Object
type ShowFileContentResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o ShowFileContentResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o ShowFileContentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFileContentResponse struct{}"
	}

	return strings.Join([]string{"ShowFileContentResponse", string(data)}, " ")
}
