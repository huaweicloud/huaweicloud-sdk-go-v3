package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadJunitCoverageZipResponse Response Object
type DownloadJunitCoverageZipResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadJunitCoverageZipResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadJunitCoverageZipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadJunitCoverageZipResponse struct{}"
	}

	return strings.Join([]string{"DownloadJunitCoverageZipResponse", string(data)}, " ")
}
