package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// CreateBlockchainCertByUserNameResponse Response Object
type CreateBlockchainCertByUserNameResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o CreateBlockchainCertByUserNameResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o CreateBlockchainCertByUserNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBlockchainCertByUserNameResponse struct{}"
	}

	return strings.Join([]string{"CreateBlockchainCertByUserNameResponse", string(data)}, " ")
}
