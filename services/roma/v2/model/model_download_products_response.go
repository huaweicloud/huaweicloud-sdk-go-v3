package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadProductsResponse Response Object
type DownloadProductsResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadProductsResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadProductsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadProductsResponse struct{}"
	}

	return strings.Join([]string{"DownloadProductsResponse", string(data)}, " ")
}
