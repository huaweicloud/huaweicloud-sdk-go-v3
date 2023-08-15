package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadBlockchainSdkConfigResponse Response Object
type DownloadBlockchainSdkConfigResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadBlockchainSdkConfigResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadBlockchainSdkConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadBlockchainSdkConfigResponse struct{}"
	}

	return strings.Join([]string{"DownloadBlockchainSdkConfigResponse", string(data)}, " ")
}
