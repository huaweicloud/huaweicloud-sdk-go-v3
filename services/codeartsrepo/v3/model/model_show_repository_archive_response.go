package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// ShowRepositoryArchiveResponse Response Object
type ShowRepositoryArchiveResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o ShowRepositoryArchiveResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o ShowRepositoryArchiveResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryArchiveResponse struct{}"
	}

	return strings.Join([]string{"ShowRepositoryArchiveResponse", string(data)}, " ")
}
