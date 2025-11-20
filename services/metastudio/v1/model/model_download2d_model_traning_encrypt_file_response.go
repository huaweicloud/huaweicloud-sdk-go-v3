package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// Download2dModelTraningEncryptFileResponse Response Object
type Download2dModelTraningEncryptFileResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o Download2dModelTraningEncryptFileResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o Download2dModelTraningEncryptFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Download2dModelTraningEncryptFileResponse struct{}"
	}

	return strings.Join([]string{"Download2dModelTraningEncryptFileResponse", string(data)}, " ")
}
