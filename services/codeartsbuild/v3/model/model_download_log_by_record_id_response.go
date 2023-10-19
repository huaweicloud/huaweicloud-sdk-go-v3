package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadLogByRecordIdResponse Response Object
type DownloadLogByRecordIdResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadLogByRecordIdResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadLogByRecordIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadLogByRecordIdResponse struct{}"
	}

	return strings.Join([]string{"DownloadLogByRecordIdResponse", string(data)}, " ")
}
