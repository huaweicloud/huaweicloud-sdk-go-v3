package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadStepImageNewResponse Response Object
type DownloadStepImageNewResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadStepImageNewResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadStepImageNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadStepImageNewResponse struct{}"
	}

	return strings.Join([]string{"DownloadStepImageNewResponse", string(data)}, " ")
}
