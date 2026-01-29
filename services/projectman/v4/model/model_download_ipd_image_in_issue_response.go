package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadIpdImageInIssueResponse Response Object
type DownloadIpdImageInIssueResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadIpdImageInIssueResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadIpdImageInIssueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadIpdImageInIssueResponse struct{}"
	}

	return strings.Join([]string{"DownloadIpdImageInIssueResponse", string(data)}, " ")
}
