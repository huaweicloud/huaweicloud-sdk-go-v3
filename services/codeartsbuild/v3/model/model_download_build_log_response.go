package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadBuildLogResponse Response Object
type DownloadBuildLogResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadBuildLogResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadBuildLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadBuildLogResponse struct{}"
	}

	return strings.Join([]string{"DownloadBuildLogResponse", string(data)}, " ")
}
