package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadImageFileResponse Response Object
type DownloadImageFileResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadImageFileResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadImageFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadImageFileResponse struct{}"
	}

	return strings.Join([]string{"DownloadImageFileResponse", string(data)}, " ")
}
