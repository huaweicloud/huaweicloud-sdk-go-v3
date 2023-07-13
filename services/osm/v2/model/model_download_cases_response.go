package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadCasesResponse Response Object
type DownloadCasesResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadCasesResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadCasesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadCasesResponse struct{}"
	}

	return strings.Join([]string{"DownloadCasesResponse", string(data)}, " ")
}
