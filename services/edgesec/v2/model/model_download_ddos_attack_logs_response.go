package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadDdosAttackLogsResponse Response Object
type DownloadDdosAttackLogsResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadDdosAttackLogsResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadDdosAttackLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadDdosAttackLogsResponse struct{}"
	}

	return strings.Join([]string{"DownloadDdosAttackLogsResponse", string(data)}, " ")
}
