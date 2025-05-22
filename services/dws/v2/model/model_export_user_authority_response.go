package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// ExportUserAuthorityResponse Response Object
type ExportUserAuthorityResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o ExportUserAuthorityResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o ExportUserAuthorityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportUserAuthorityResponse struct{}"
	}

	return strings.Join([]string{"ExportUserAuthorityResponse", string(data)}, " ")
}
