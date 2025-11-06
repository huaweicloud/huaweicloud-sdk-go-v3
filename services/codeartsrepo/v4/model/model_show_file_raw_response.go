package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// ShowFileRawResponse Response Object
type ShowFileRawResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o ShowFileRawResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o ShowFileRawResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFileRawResponse struct{}"
	}

	return strings.Join([]string{"ShowFileRawResponse", string(data)}, " ")
}
