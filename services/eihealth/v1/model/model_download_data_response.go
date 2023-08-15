package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadDataResponse Response Object
type DownloadDataResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadDataResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadDataResponse struct{}"
	}

	return strings.Join([]string{"DownloadDataResponse", string(data)}, " ")
}
