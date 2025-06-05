package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadKeystoreByNameResponse Response Object
type DownloadKeystoreByNameResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadKeystoreByNameResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadKeystoreByNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadKeystoreByNameResponse struct{}"
	}

	return strings.Join([]string{"DownloadKeystoreByNameResponse", string(data)}, " ")
}
