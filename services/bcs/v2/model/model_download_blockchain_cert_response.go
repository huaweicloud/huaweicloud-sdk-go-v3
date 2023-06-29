package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadBlockchainCertResponse Response Object
type DownloadBlockchainCertResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadBlockchainCertResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadBlockchainCertResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadBlockchainCertResponse struct{}"
	}

	return strings.Join([]string{"DownloadBlockchainCertResponse", string(data)}, " ")
}
