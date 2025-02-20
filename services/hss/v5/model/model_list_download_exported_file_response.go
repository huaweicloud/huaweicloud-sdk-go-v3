package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// ListDownloadExportedFileResponse Response Object
type ListDownloadExportedFileResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o ListDownloadExportedFileResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o ListDownloadExportedFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDownloadExportedFileResponse struct{}"
	}

	return strings.Join([]string{"ListDownloadExportedFileResponse", string(data)}, " ")
}
