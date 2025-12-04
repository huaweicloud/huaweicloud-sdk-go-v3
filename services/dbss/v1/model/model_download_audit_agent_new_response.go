package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadAuditAgentNewResponse Response Object
type DownloadAuditAgentNewResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadAuditAgentNewResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadAuditAgentNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadAuditAgentNewResponse struct{}"
	}

	return strings.Join([]string{"DownloadAuditAgentNewResponse", string(data)}, " ")
}
