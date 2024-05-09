package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadCompareResultFileResponse Response Object
type DownloadCompareResultFileResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadCompareResultFileResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadCompareResultFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadCompareResultFileResponse struct{}"
	}

	return strings.Join([]string{"DownloadCompareResultFileResponse", string(data)}, " ")
}
