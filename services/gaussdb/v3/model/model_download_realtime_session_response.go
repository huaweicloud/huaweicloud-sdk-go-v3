package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadRealtimeSessionResponse Response Object
type DownloadRealtimeSessionResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadRealtimeSessionResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadRealtimeSessionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadRealtimeSessionResponse struct{}"
	}

	return strings.Join([]string{"DownloadRealtimeSessionResponse", string(data)}, " ")
}
