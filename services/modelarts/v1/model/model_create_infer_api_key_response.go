package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// CreateInferApiKeyResponse Response Object
type CreateInferApiKeyResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o CreateInferApiKeyResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o CreateInferApiKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInferApiKeyResponse struct{}"
	}

	return strings.Join([]string{"CreateInferApiKeyResponse", string(data)}, " ")
}
