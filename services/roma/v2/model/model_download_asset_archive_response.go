package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadAssetArchiveResponse Response Object
type DownloadAssetArchiveResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadAssetArchiveResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadAssetArchiveResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadAssetArchiveResponse struct{}"
	}

	return strings.Join([]string{"DownloadAssetArchiveResponse", string(data)}, " ")
}
